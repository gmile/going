package main

import (
  client "ttt_client"
)

func main() {
  game := new(client.Game)
  game.Join() // should be game.Join('127.0.0.1', ':1234')
}
