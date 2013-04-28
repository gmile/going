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

func (game Game) Start() {
	server, _ := net.Listen("tcp", ":1234")
  defer server.Close()

  fmt.Printf("Waiting for other player to connect...")

  conn, _ := server.Accept()
  defer conn.Close()

  fmt.Printf(" opponent connected!\n")

  game.Conn = conn

  settings := bufio.NewWriter(game.Conn)
  game.set_marks(settings)
  game.set_turn_sequence(settings)
  settings.Flush()

  game.Play()
}

func (game Game) set_turn_sequence(writer *bufio.Writer) {
  game.I_go_first = byte(rand.Intn(2))
  writer.WriteByte(1 - game.I_go_first)
}

func (game Game) set_marks(writer *bufio.Writer) {
  if rand.Intn(2) == 0 {
    game.Player1_mark = 'X'
    game.Player2_mark = 'O'
  } else {
    game.Player1_mark = 'O'
    game.Player2_mark = 'X'
  }

  writer.WriteRune(game.Player2_mark)
  writer.WriteRune(game.Player1_mark)
}
