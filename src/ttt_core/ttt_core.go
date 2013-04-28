package ttt_core

import (
	"bufio"
	"fmt"
	"net"
)

// player1 -> current_player
// player2 -> other_player
// struct Player {
//   turn_order byte
//   mark rune
// }

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

func (game *Game) recordTurn(mark rune, row rune, col byte) bool {
	var y byte = 0
	var x byte = col - 1

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
  return game.checkConditions(mark, x, y)
}

func (game *Game) checkConditions(mark rune, x, y byte) (win bool) {
  b := game.board

  if b[0][x] | b[1][x] | b[2][x] == mark ||
     b[y][0] | b[y][1] | b[y][2] == mark ||
     b[0][0] | b[1][1] | b[2][2] == mark ||
     b[0][2] | b[1][1] | b[2][0] == mark {
    win = true
  }

  return
}

func (game *Game) readInput() (col byte, row rune) {
  fmt.Scanf("%c%d", &row, &col)
  return
}

func (game *Game) makeTurn() (win bool) {
  fmt.Printf("\nMake a turn: ")

  col, row := game.readInput()

  win = game.recordTurn(game.Player1Mark, row, col)

  var win_code byte = 0
  if win {
    win_code = 1
  } else {
    win_code = 0
  }

  if win {
    fmt.Printf("\nYou won!\n")
  }

  my_turn := bufio.NewWriter(game.Conn)
  my_turn.WriteRune(row)
  my_turn.WriteByte(col)
  my_turn.WriteByte(win_code)
  my_turn.Flush()

  return
  // TODO: game.notifySpectators(x, y, mark)
}

func (game *Game) waitForOtherTurn() (win bool) {
  fmt.Printf("\nWaiting for opponent...\n")

  his_turn := bufio.NewReader(game.Conn)

  row, _, _   := his_turn.ReadRune()
  col, _      := his_turn.ReadByte()
  win_code, _ := his_turn.ReadByte()

  game.recordTurn(game.Player2Mark, row, col)

  win = win_code == 1

  if win {
    fmt.Printf("\nOther player won.\n")
  }

  return
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
      if game.makeTurn() { break }
      if game.waitForOtherTurn() { break }
    }
  } else {
    fmt.Printf("\nOpponent's turn is first.\n")

    for {
      if game.waitForOtherTurn() { break }
      if game.makeTurn() { break }
    }
  }
}
