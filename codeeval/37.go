package main

import (
  "os"
  "bufio"
)

func main() {
  file, _ := os.Open( os.Args[1] )
  scanner := bufio.NewScanner(file)

  var alphabet uint32 = 0x3FFFFFF

  for scanner.Scan() {
    var result uint32 = 0

    line := scanner.Text()

    if len(line) == 0 {
      continue
    }

    for _, r := range line {
      var c byte

      if r >= 65 && r <= 90 {
        c = byte(r) - 65
      } else if r >= 97 && r <= 122 {
        c = byte(r) - 97
      }

      result |= (1 << c)
    }

    if result == alphabet {
      print("NULL")
    } else {
      var c uint32 = 1

      for i := 0; i < 26; i++ {
        if (result & c) == 0 {
          print(string(i + 97))
        }

        c <<= 1
      }
    }

    println()
  }
}
