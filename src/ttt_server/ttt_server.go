package ttt_server

import (
  core "ttt_core"
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
  game.Player1Turn = byte(rand.Intn(2))
  game.Player2Turn = 1 - game.Player1Turn

  game.SendBuffer.WriteByte(game.Player2Turn)
  game.SendBuffer.WriteByte(game.Player1Turn)
}

func (game *Game) setMarks() {
  if rand.Intn(2) == 0 {
    game.Player1Mark = 'X'
    game.Player2Mark = 'O'
  } else {
    game.Player1Mark = 'O'
    game.Player2Mark = 'X'
  }

  game.SendBuffer.WriteRune(game.Player2Mark)
  game.SendBuffer.WriteRune(game.Player1Mark)
}
