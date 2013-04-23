package main

import "net"
import "bufio"

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

    reader := bufio.NewReader(conn)
    my_string, err := reader.ReadString('\n')

    if err != nil {
      println("You said:", my_string)
    }
  }
}
