package ttt_client

import (
  core "ttt_core"
	"bufio"
	"net"
)

type Game struct {
  core.Game
}

func (game *Game) Join() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  game.Conn = conn

  settings := bufio.NewReader(game.Conn)
  game.getMarks(settings)
  game.getTurnSequence(settings)

  game.Play()
}

func (game *Game) getTurnSequence(reader *bufio.Reader) {
  game.Player1Turn, _ = reader.ReadByte()
  game.Player2Turn, _ = reader.ReadByte()
}

func (game *Game) getMarks(reader *bufio.Reader) {
  game.Player1Mark, _, _ = reader.ReadRune()
  game.Player2Mark, _, _ = reader.ReadRune()
}
