package server

import (
  core "ttt"
	"bufio"
	"fmt"
	"net"
  "math/rand"
)

type Game struct {
  core.Game 
}

func (game *Game) Start() {
	server, _ := net.Listen("tcp", ":1234")
  defer server.Close()

  fmt.Printf("Waiting for other player to connect...")

  conn, _ := server.Accept()
  defer conn.Close()

  fmt.Printf(" connected!\n")

  game.SendBuffer    = bufio.NewWriter(conn)
  game.ReceiveBuffer = bufio.NewReader(conn)
  game.setMarks()
  game.setTurnSequence()
  game.SendBuffer.Flush()

  game.Play()
}

func (game *Game) setTurnSequence() {
  game.CurrentPlayer.TurnOrder = byte(rand.Intn(2))
  game.OtherPlayer.TurnOrder = 1 - game.CurrentPlayer.TurnOrder

  game.SendBuffer.WriteByte(game.OtherPlayer.TurnOrder)
  game.SendBuffer.WriteByte(game.CurrentPlayer.TurnOrder)
}

func (game *Game) setMarks() {
  if rand.Intn(2) == 0 {
    game.CurrentPlayer.Mark = 'X'
    game.OtherPlayer.Mark = 'O'
  } else {
    game.CurrentPlayer.Mark = 'O'
    game.OtherPlayer.Mark = 'X'
  }

  game.SendBuffer.WriteRune(game.OtherPlayer.Mark)
  game.SendBuffer.WriteRune(game.CurrentPlayer.Mark)
}
