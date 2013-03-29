package main

import (
  "log"
)

func main() {
  sum := 0

  for i := 1; i < 1000; i++ {
    if (i % 3 == 0) || (i % 5 == 0) {
      sum += i
    }
  }

  log.Println(sum)
}
