package ttt_core

import (
	"bufio"
	"fmt"
	"net"
)

type Game struct {
  Player1_mark rune
  Player2_mark rune
  I_go_first byte
  Conn net.Conn
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

  // TODO: notify spectators here
}

func (game *Game) readInput() (col byte, row rune) {
  fmt.Scanf("%c%d", &row, &col)
  return
}

func (game *Game) MakeTurn() {
  fmt.Printf("\nMake a turn: ")

  col, row := game.readInput()

  my_turn := bufio.NewWriter(game.Conn)
  my_turn.WriteRune(row)
  my_turn.WriteByte(col)
  my_turn.Flush()

  game.RecordTurn(game.Player1_mark, row, col)
}

func (game *Game) WaitForOtherTurn() {
  fmt.Printf("\nWaiting for opponent...\n")

  his_turn := bufio.NewReader(game.Conn)

  row, _, _ := his_turn.ReadRune()
  col, _    := his_turn.ReadByte()

  game.RecordTurn(game.Player2_mark, row, col)
}

func (game *Game) Play() {
  game.board = [3][3]rune{
    {' ', ' ', ' '},
    {' ', ' ', ' '},
    {' ', ' ', ' '}}

  game.DrawBoard()

  if game.I_go_first == 1 {
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
