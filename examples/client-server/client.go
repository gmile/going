package main

import "net"
import "bufio"

func main() {
  conn, err := net.Dial("tcp", ":8080")

  if err != nil {
    println("There was an error:", err)
  }

  my_string_writer := bufio.NewWriter(conn)

  _, err = my_string_writer.WriteString("Example")
  err = my_string_writer.Flush()
}
