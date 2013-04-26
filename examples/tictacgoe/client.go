package main

import (
  "net"
  "bufio"
)

func main() {
  conn, _ := net.Dial("tcp", ":1234")
  defer conn.Close()

  writer := bufio.NewWriter(conn)
  defer writer.Flush()

  writer.WriteRune('a')
  writer.WriteByte(1)
}
