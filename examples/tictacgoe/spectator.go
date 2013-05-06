package main

import (
  "ttt/spectator"
)

func main() {
  game := new(spectator.Game)
  game.Spectate() // TODO: should be game.Join('127.0.0.1', ':1234')
}
