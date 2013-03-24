package main

import (
  "fmt" 
  "strconv"
)

func main() {
  fmt.Println("Example 1")
  hashmap := map[string]int{}
  key := ""

  for i := 0; i < 4; i++ {
    key = strconv.Itoa(i)
    hashmap[key + " -"] = i+15
    fmt.Println(hashmap[key])
  }

  fmt.Println("Example 2")
  not_exists, ok := hashmap["example"]
  fmt.Println(not_exists, ok)

  fmt.Println("Example 3")

  for k, v := range hashmap {
    fmt.Println(k, v)
  }
}
