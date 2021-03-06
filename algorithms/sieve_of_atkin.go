package main

import (
  m "math" 
)

const LIMIT = 10000

var (
  n, q int
)

func main() {
  prime_bools := [LIMIT + 1]bool{}

  for i := 0; i < 5; i++ {
    prime_bools[i] = true
  }

  sqrt := int( m.Ceil(m.Sqrt(LIMIT)) )

  field := make([][]int, sqrt)
  for i, _ := range field {
    field[i] = make([]int, sqrt)
  }


  for i := 1; i <= sqrt; i++ {
    for j := 1; j <= sqrt; j++ {
      n = 4 * i * i + j * j
      if (n <= LIMIT) && (n % 12 == 1 || n % 12 == 5) { prime_bools[n] = !prime_bools[n] }

      n = 3 * i * i + j * j
      if (n <= LIMIT) && (n % 12 == 7)                { prime_bools[n] = !prime_bools[n] }

      n = 3 * i * i - j * j
      if (i > j) && (n <= LIMIT) && (n % 12 == 11)    { prime_bools[n] = !prime_bools[n] }
    }
  }

  for i := 5; i <= LIMIT; i++ {
    if prime_bools[i] {
      q = i * i

      for k := 1; k*q <= LIMIT; k++ {
        prime_bools[k*q] = false
      }
    }
  }

  println(2)
  println(3)

  for i := 5; i <= LIMIT; i++ {
    if prime_bools[i] { println(i) }
  }
}
