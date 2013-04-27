package main

import (
  "net"
  //"fmt"
  "bufio"
)

func main() {
  conn, _ := net.Dial("tcp", ":1234")

  writer := bufio.NewWriter(conn)

  writer.WriteRune('b')
  writer.WriteByte(1)
  writer.Flush()

  conn.Close()
}
