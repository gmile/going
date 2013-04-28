package ttt_core

import (
	"bufio"
	"fmt"
	"net"
)

type Game struct {
  Player1Mark rune
  Player2Mark rune
  Player1Turn byte
  Player2Turn byte
  Conn net.Conn
  board [3][3]rune
}

func (game *Game) DrawBoard() {
	fmt.Printf("\n    1   2   3\n")
	fmt.Printf(" a [%c] [%c] [%c]\n", game.board[0][0], game.board[0][1], game.board[0][2])
	fmt.Printf(" b [%c] [%c] [%c]\n", game.board[1][0], game.board[1][1], game.board[1][2])
	fmt.Printf(" c [%c] [%c] [%c]\n", game.board[2][0], game.board[2][1], game.board[2][2])
}

func (game *Game) recordTurn(mark rune, row rune, col byte) {
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

func (game *Game) makeTurn() {
  fmt.Printf("\nMake a turn: ")

  col, row := game.readInput()

  my_turn := bufio.NewWriter(game.Conn)
  my_turn.WriteRune(row)
  my_turn.WriteByte(col)
  my_turn.Flush()

  game.recordTurn(game.Player1Mark, row, col)
}

func (game *Game) waitForOtherTurn() {
  fmt.Printf("\nWaiting for opponent...\n")

  his_turn := bufio.NewReader(game.Conn)

  row, _, _ := his_turn.ReadRune()
  col, _    := his_turn.ReadByte()

  game.recordTurn(game.Player2Mark, row, col)
}

func (game *Game) Play() {
  game.board = [3][3]rune{
    {' ', ' ', ' '},
    {' ', ' ', ' '},
    {' ', ' ', ' '}}

  game.DrawBoard()

  if game.Player1Turn == 1 {
    fmt.Printf("\nYour turn is first.\n")

    for {
      game.makeTurn()
      game.waitForOtherTurn()
    }
  } else {
    fmt.Printf("\nOpponent's turn is first.\n")

    for {
      game.waitForOtherTurn()
      game.makeTurn()
    }
  }
}
