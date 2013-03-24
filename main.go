package main

import (
  log "fmt"
  i "ingestion"
)

func main() {
  var inventory i.Inventory
  inventory.Initialize("/Users/gmile/odesk/haystack_profiling/homenet.txt")

  for _, car := range inventory.SingleCars() {
    log.Println(&car, car.Model, "was made in", car.Year)
  }
}
