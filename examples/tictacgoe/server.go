package main

import (
  server "ttt_server"
)

func main() {
  game := new(server.Game)
  game.Start() // should be game.Start('127.0.0.1', ':1234')
}
