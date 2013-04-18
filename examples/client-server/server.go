package main

import "net"

func main() {
  ln, err := net.Listen("tcp", ":8080")

  if err != nil {
    println("There was an error:", err)
  }

  for {
    conn, err := ln.Accept()

    if err != nil {
      println("Didn't work")
    }

    buffer := make([]byte, 2)
    size, err := conn.Read(buffer)

    if err != nil {
      println("Didn't work")
    }

    println(size)
    println(buffer[0])
    println(buffer[1])
  }
}
