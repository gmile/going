package main

import "fmt"

func print_matrix(m *[3][3]rune) {
  fmt.Printf("\n    1   2   3\n")
  fmt.Printf(" a [%c] [%c] [%c]\n", m[0][0], m[0][1], m[0][2])
  fmt.Printf(" b [%c] [%c] [%c]\n", m[1][0], m[1][1], m[1][2])
  fmt.Printf(" c [%c] [%c] [%c]\n", m[2][0], m[2][1], m[2][2])
}

func main() {
  m := [3][3]rune {
    {' ', ' ', ' '},
    {' ', ' ', ' '},
    {' ', ' ', ' '} }

  print_matrix(&m)

  received_row  := 'b'
  received_col  := 1
  received_mark := 'X'

  y := received_row
  x := received_col - 1

  switch y {
  case 'a': y = 0
  case 'b': y = 1
  case 'c': y = 2
  }

  m[y][x] = received_mark
  fmt.Printf("\n Waiting for opponent...")
  fmt.Printf("\n[%d, %d] = %c\n", x, y, received_mark)

  print_matrix(&m)
}
