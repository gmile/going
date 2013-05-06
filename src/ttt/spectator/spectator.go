package spectator

import (
  core "ttt"
	"bufio"
	"net"
	"fmt"
)

type Game struct {
  core.Game
}

func (game *Game) Spectate() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  reader := bufio.Reader

  var col, win_code byte
  var row, mark     rune

  for {
    println("Waiting for someone to make a turn...")
    col      = reader.ReadByte()
    row      = reader.ReadRune()
    mark     = reader.ReadRune()
    win_code = reader.ReadByte()

    println(col, row, mark, win_code)
  }
}
