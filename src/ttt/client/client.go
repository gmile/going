package client

import (
  core "ttt"
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
  game.CurrentPlayer.TurnOrder, _ = game.ReceiveBuffer.ReadByte()
  game.OtherPlayer.TurnOrder, _ = game.ReceiveBuffer.ReadByte()
}

func (game *Game) getMarks() {
  game.CurrentPlayer.Mark, _, _ = game.ReceiveBuffer.ReadRune()
  game.OtherPlayer.Mark, _, _ = game.ReceiveBuffer.ReadRune()
}
