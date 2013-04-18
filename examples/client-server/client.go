package main

import "net"

func main() {
  conn, err := net.Dial("tcp", ":8080")

  if err != nil {
    println("There was an error:", err)
  }

  buffer := make([]byte, 2)
  buffer[0] = 1
  buffer[1] = 251

  conn.Write(buffer)
}
