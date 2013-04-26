package main

import "net"
import "net/rpc"

type Region struct {
  X, Y int
}

type MyRegionServer int

func (s *MyRegionServer) GetRegion(id *int, reply *Region) error {
  println("Someone want's a region with id", *id)

  found_region := Region{10,20}
  reply = &found_region

  println("Let's give him a region", reply.X, reply.Y)

  return nil
}

func main() {
  listener, _ := net.Listen("tcp", ":8080")

  rpc_server := rpc.NewServer()
  rpc_server.Register( new(MyRegionServer) )

  for {
    rpc_server.Accept(listener)
  }
}
