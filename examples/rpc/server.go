package main

import "tcp/rpc"

type Region struct {
  X, Y int
}

type Server int

func (s *Server) GetRegion(id *int, reply *Region) {
  // 1. find the region in DB
  // 2. return it back

  found_region = new(Region)
  found_region.X = 10
  found_region.Y = 20

  *Region = found_region

  return nil
}

func main() {
  server := new(Server)

  rpc.Register(server)
  // rpc.HandleHTTP()

  l, e := net.Listen("tcp", ":1234")

  if e != nil {
    log.Fatal("listen error:", e)
  }

  go http.Serve(l, nil)
}
