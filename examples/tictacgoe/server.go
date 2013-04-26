package main

import (
	"bufio"
	"fmt"
	"net"
)

func print_matrix(m *[3][3]rune) {
	fmt.Printf("\n    1   2   3\n")
	fmt.Printf(" a [%c] [%c] [%c]\n", m[0][0], m[0][1], m[0][2])
	fmt.Printf(" b [%c] [%c] [%c]\n", m[1][0], m[1][1], m[1][2])
	fmt.Printf(" c [%c] [%c] [%c]\n", m[2][0], m[2][1], m[2][2])
}

func turn(m *[3][3]rune, mark rune, x, y byte) {
	(*m)[y][x] = mark
}

func opponents_turn(m *[3][3]rune, mark rune, row rune, col byte) {
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

	(*m)[y][x] = mark
}

func main() {
	m := [3][3]rune{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '}}

	listener, _ := net.Listen("tcp", ":1234")

	fmt.Printf("\n Waiting for opponent...\n")

  // my_mark
  // opponents_mark

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		reader := bufio.NewReader(conn)

		received_rune := 'X'
		received_row, _, _ := reader.ReadRune()
		received_col, _ := reader.ReadByte()

		opponents_turn(&m, received_rune, received_row, received_col)
		print_matrix(&m)
	}
}
