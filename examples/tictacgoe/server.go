package main

import (
	"bufio"
	"fmt"
	"net"
)

type Game struct {
  player1_mark rune
  player2_mark rune
  game_starts_from byte
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
}

func (game *Game) Player1Turn() {
  fmt.Printf("\nMake a turn:\n")

  // TODO: ask for input
  var col byte = 1
  row := 'a'

  my_turn := bufio.NewWriter(game.conn)
  defer my_turn.Flush()

  my_turn.WriteRune(row)
  my_turn.WriteByte(col)

  game.RecordTurn(game.player1_mark, row, col)
}

func (game *Game) Player2Turn() {
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

  game.set_marks()
  game.set_turn_sequence()
  game.play()
}

func (game *Game) Join() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  game.conn = conn

  game.get_marks()
  game.get_turn_sequence()
  game.play()
}

func (game *Game) set_turn_sequence() {
  // TODO: pick this randomly from (0, 1)
  game.game_starts_from = 2

  writer := bufio.NewWriter(game.conn)
  writer.WriteByte(game.game_starts_from)
  writer.Flush()
}

func (game *Game) get_turn_sequence() {
  reader := bufio.NewReader(game.conn)
  game.game_starts_from, _ = reader.ReadByte()
}

func (game *Game) set_marks() {
  // TODO: pick these randomly
  player1_mark := 'X'
  player2_mark := 'O'

  game.player1_mark = player1_mark
  game.player2_mark = player2_mark

  marks := bufio.NewWriter(game.conn)
  defer marks.Flush()

  marks.WriteRune(player1_mark)
  marks.WriteRune(player2_mark)
}

func (game *Game) get_marks() {
  marks := bufio.NewReader(game.conn)

  game.player1_mark, _, _ = marks.ReadRune()
  game.player2_mark, _, _ = marks.ReadRune()
}

func (game *Game) play() {
  game.board = [3][3]rune{
    {' ', ' ', ' '},
    {' ', ' ', ' '},
    {' ', ' ', ' '}}

  game.DrawBoard()

  if game.game_starts_from == 1 {
    fmt.Printf("Your turn is first.\n")

    for {
      game.Player1Turn()
      game.Player2Turn()
    }
  } else {
    fmt.Printf("Opponent's turn is first.\n")

    for {
      game.Player2Turn()
      game.Player1Turn()
    }
  }
}

func main() {
  game := new(Game)
  game.Start() // should be game.Start('127.0.0.1', ':1234')
}
