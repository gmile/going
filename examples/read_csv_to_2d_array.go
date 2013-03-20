package main

// Reads CSV from stdin and dumps it back to stdout.

import (
  "code.google.com/p/gocsv"
  "os"
  "fmt"
)

// parses a csv from stdin
func parse_csv() [][]string {
  rows, err := csv.ReadAll(os.Stdin)

  if err != nil {
    panic(err)
  }

  return rows
}

func main() {
  rows := parse_csv()

  for _, row := range rows[1:2] {
    for _, cell := range row {
      fmt.Println(cell)
    }
  }
}
