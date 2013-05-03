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

  game.SendBuffer    = bufio.NewWriter(conn)
  game.ReceiveBuffer = bufio.NewReader(conn)
  game.getMarks()
  game.getTurnSequence()

  game.Play()
}

func (game *Game) getTurnSequence() {
  game.Player1Turn, _ = game.ReceiveBuffer.ReadByte()
  game.Player2Turn, _ = game.ReceiveBuffer.ReadByte()
}

func (game *Game) getMarks() {
  game.Player1Mark, _, _ = game.ReceiveBuffer.ReadRune()
  game.Player2Mark, _, _ = game.ReceiveBuffer.ReadRune()
}
