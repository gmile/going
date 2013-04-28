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

  game.Conn = conn

  settings := bufio.NewWriter(game.Conn)
  game.setMarks(settings)
  game.setTurnSequence(settings)
  settings.Flush()

  game.Play()
}

func (game *Game) setTurnSequence(writer *bufio.Writer) {
  game.I_go_first = byte(rand.Intn(2))
  writer.WriteByte(1 - game.I_go_first)
}

func (game *Game) setMarks(writer *bufio.Writer) {
  if rand.Intn(2) == 0 {
    game.Player1Mark = 'X'
    game.Player2Mark = 'O'
  } else {
    game.Player1Mark = 'O'
    game.Player2Mark = 'X'
  }

  writer.WriteRune(game.Player2Mark)
  writer.WriteRune(game.Player1Mark)
}