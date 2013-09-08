package main

import (
  "os"
  "bufio"
  "strings"
)

func main() {
  file, _ := os.Open( os.Args[1] )
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()

    if len(line) > 0 {
      values     := strings.Split(line, ",")
      str        := values[0]
      symbol     := values[1]
      last_index := strings.LastIndex(str, symbol)

      println(last_index)
    }
  }
}
