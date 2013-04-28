package main

import (
	"bufio"
	"fmt"
	"net"
  "math/rand"
)

// refactor to my_mark and my_turn?

type Game struct {
  player1_mark rune
  player2_mark rune
  i_go_first byte
  conn net.Conn
  board [3][3]rune
}

func (game *Game) DrawBoard() {
	fmt.Printf("\n    1   2   3\n")
	fmt.Printf(" a [%c] [%c] [%c]\n", game.board[0][0], game.board[0][1], game.board[0][2])
	fmt.Printf(" b [%c] [%c] [%c]\n", game.board[1][0], game.board[1][1], game.board[1][2])
	fmt.Printf(" c [%c] [%c] [%c]\n", game.board[2][0], game.board[2][1], game.board[2][2])
}

// TODO: Split this into GameCore, GameClient and GameServer

func (game *Game) RecordTurn(mark rune, row rune, col byte) {
	y := 0
	x := col - 1

	switch row {
	case 'a':
		y = 0
	case 'b':
		y = 1
	case 'c':
		y = 2
	}

	game.board[y][x] = mark

  game.DrawBoard()

  // TODO: notify spectators
}

func (game *Game) readInput() (col byte, row rune) {
  fmt.Scanf("%c%d", &row, &col)
  return
}

func (game *Game) MakeTurn() {
  fmt.Printf("\nMake a turn: ")

  col, row := game.readInput()

  my_turn := bufio.NewWriter(game.conn)
  my_turn.WriteRune(row)
  my_turn.WriteByte(col)
  my_turn.Flush()

  game.RecordTurn(game.player1_mark, row, col)
}

func (game *Game) WaitForOtherTurn() {
  fmt.Printf("\nWaiting for opponent...\n")

  his_turn := bufio.NewReader(game.conn)

  row, _, _ := his_turn.ReadRune()
  col, _    := his_turn.ReadByte()

  game.RecordTurn(game.player2_mark, row, col)
}

func (game *Game) Start() {
	server, _ := net.Listen("tcp", ":1234")
  defer server.Close()

  fmt.Printf("Waiting for other player to connect...")

  conn, _ := server.Accept()
  defer conn.Close()

  fmt.Printf(" opponent connected!\n")

  game.conn = conn

  writer := bufio.NewWriter(game.conn)
  game.set_marks(writer)
  game.set_turn_sequence(writer)
  writer.Flush()

  game.play()
}

func (game *Game) Join() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  game.conn = conn

  reader := bufio.NewReader(game.conn)
  game.get_marks(reader)
  game.get_turn_sequence(reader)

  game.play()
}

func (game *Game) set_turn_sequence(writer *bufio.Writer) {
  game.i_go_first = byte(rand.Intn(2))
  writer.WriteByte(1 - game.i_go_first)
}

func (game *Game) get_turn_sequence(reader *bufio.Reader) {
  game.i_go_first, _ = reader.ReadByte()
}

func (game *Game) set_marks(writer *bufio.Writer) {
  if rand.Intn(2) == 0 {
    game.player1_mark = 'X'
    game.player2_mark = 'O'
  } else {
    game.player1_mark = 'O'
    game.player2_mark = 'X'
  }

  writer.WriteRune(game.player2_mark)
  writer.WriteRune(game.player1_mark)
}

func (game *Game) get_marks(reader *bufio.Reader) {
  game.player1_mark, _, _ = reader.ReadRune()
  game.player2_mark, _, _ = reader.ReadRune()
}

func (game *Game) play() {
  game.board = [3][3]rune{
    {' ', ' ', ' '},
    {' ', ' ', ' '},
    {' ', ' ', ' '}}

  game.DrawBoard()

  if game.i_go_first == 1 {
    fmt.Printf("\nYour turn is first.\n")

    for {
      game.MakeTurn()
      game.WaitForOtherTurn()
    }
  } else {
    fmt.Printf("\nOpponent's turn is first.\n")

    for {
      game.WaitForOtherTurn()
      game.MakeTurn()
    }
  }
}

func main() {
  game := new(Game)
  game.Join() // should be game.Start('127.0.0.1', ':1234')
}
