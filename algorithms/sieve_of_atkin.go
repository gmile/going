package main

import (
  m "math" 
)

const LIMIT = 10000
var int n

func main() {
  prime_bools := [LIMIT + 1]bool{}

  for i := 0; i < 5; i++ {
    prime_bools := true
  }

  sqrt = m.Sqrt(LIMIT)

  field := make(int[][], sqrt)
  for i, _ := range field {
    field[i] = make(int[], sqrt)
  }

  for i := 1; i <= sqrt; i++ {
    for j := 1; j <= sqrt; i++ {
      n = 4 * m.Pow(i, 2) + m.Pow(j, 2)
      if (n <= LIMIT) && (n % 12 == 1 || n % 12 == 5) { prime_bools[n] = !prime_bools[n] }

      n = 3 * m.Pow(i, 2) + m.Pow(j, 2)
      if (n <= LIMIT) && (n % 12 == 7)                { prime_bools[n] = !prime_bools[n] }

      n = 4 * m.Pow(i, 2) - m.Pow(j, 2)
      if (i > j) && (n <= LIMIT) && (n % 12 == 11)    { prime_bools[n] = !prime_bools[n] }
    }
  }

  for i := 5; i <= LIMIT; i++ {
    if prime_bools[i] {
      q = m.Pow(i, 2)

      k := 1; k <= LIMIT; k++ {
        prime_bools[k*q] = false
      }
    }
  }

  print(2, 3)

  for i := 5; i <= LIMIT; i++ {
    if prime_bools { print i }
  }
}
