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
  game.get_marks(settings)
  game.get_turn_sequence(settings)

  game.Play()
}

func (game *Game) get_turn_sequence(reader *bufio.Reader) {
  game.I_go_first, _ = reader.ReadByte()
}

func (game *Game) get_marks(reader *bufio.Reader) {
  game.Player1_mark, _, _ = reader.ReadRune()
  game.Player2_mark, _, _ = reader.ReadRune()
}
