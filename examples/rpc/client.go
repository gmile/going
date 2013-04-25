package main

import "net/rpc"

type Region struct {
  X, Y int
}

func main() {
  client, _ := rpc.Dial("tcp", ":8080")

  x := Region{0,0}

  div_call := client.Go("Server.GetRegion", 1, &x, nil)
  <-div_call.Done

  println(x.X, x.Y)
}
