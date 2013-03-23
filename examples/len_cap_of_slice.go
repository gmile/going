package main

import "fmt"

func main() {
  x := make( []byte, 5)

  inspectSlice(x)

  for i := 0; i < 4; i++ {
    x = append(x, 123)
  }

  inspectSlice(x)
}

func inspectSlice(slice []byte) {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("work failed:", err)
    }
  }()

  fmt.Println( "len(x) =", len(slice) )
  fmt.Println( "cap(x) =", cap(slice) )
  fmt.Println( "x[5]   =", slice[5] )
}
