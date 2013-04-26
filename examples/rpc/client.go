package main

import "net/rpc"

type Region struct {
  X, Y int
}

func main() {
  client, _ := rpc.Dial("tcp", ":8080")

  region := Region{0,0}

  div_call := client.Go("MyRegionServer.GetRegion", 1, &region, nil)
  <-div_call.Done

  println(region.X, region.Y)
}
