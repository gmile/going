package main

import "tcp/rpc"

type Region struct {
  X, Y int
}

func main() {
  client, err := rpc.Dial("tcp", "localhost:1234")

  if err != nil {
    log.Fatal("dialing:", err)
  }
}
