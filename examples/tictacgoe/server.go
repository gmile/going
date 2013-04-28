package main

import (
  core "ttt_core"
  core "ttt_server"
	"bufio"
	"fmt"
	"net"
  "math/rand"
)

type ServerGame struct {
  core.Game
}

type ClientGame struct {
  core.Game
}

func (game ServerGame) Start() {
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

func (game ClientGame) Join() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  game.Conn = conn

  reader := bufio.NewReader(game.Conn)
  game.get_marks(reader)
  game.get_turn_sequence(reader)

  game.Play()
}

func (game ServerGame) set_turn_sequence(writer *bufio.Writer) {
  game.I_go_first = byte(rand.Intn(2))
  writer.WriteByte(1 - game.I_go_first)
}

func (game ClientGame) get_turn_sequence(reader *bufio.Reader) {
  game.I_go_first, _ = reader.ReadByte()
}

func (game ServerGame) set_marks(writer *bufio.Writer) {
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

func (game ClientGame) get_marks(reader *bufio.Reader) {
  game.Player1_mark, _, _ = reader.ReadRune()
  game.Player2_mark, _, _ = reader.ReadRune()
}

func main() {
  game := new(ServerGame)
  game.Start() // should be game.Start('127.0.0.1', ':1234')
}
