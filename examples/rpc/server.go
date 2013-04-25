package main

import "net"
import "net/rpc"

type Region struct {
  X, Y int
}

type Server int

func (s *Server) GetRegion(id *int, reply *Region) error {
  reply.X = 1
  reply.Y = 1

  return nil
}

func main() {
  listener, _ := net.Listen("tcp", ":8080")

  rpc_server := rpc.NewServer()
  rpc_server.Register( new(Server) )

  for {
    rpc_server.Accept(listener)
  }
}
