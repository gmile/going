package main

import "fmt"

func main() {
  for pos, char := range "aΦz" {
    fmt.Printf("%c - %d\n", char, pos)
  }
}
