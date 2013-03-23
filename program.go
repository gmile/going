package main

import (
  "fmt"
  i "inventory"
)

func main() {
  var inventory i.Inventory
  inventory.Initialize("/Users/gmile/odesk/haystack_profiling/homenet.txt")

  for _, car := range inventory.Cars[0:5] {
    fmt.Println(car.Model, "was made in", car.Year)
  }
}
