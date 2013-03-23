package main

import (
  "fmt" 
  "strconv"
)

func main() {
  hashmap := map[string]int{}
  key := ""

  for i := 0; i < 4; i++ {
    key = strconv.Itoa(i)
    hashmap[key] = i+15
    fmt.Println(hashmap[key])
  }
}
