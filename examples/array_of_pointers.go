package main

import (
  log "fmt" 
)

func returnSelected(incomingNumbers []string) (selectedObjects []*string) {
  selected := make([]*string, 3)

  selected[0] = &incomingNumbers[0]
  selected[1] = &incomingNumbers[5]
  selected[2] = &incomingNumbers[9]

  return selected
}

func main() {
  my_array := make([]string, 10)
  indexes  := []int{ 0, 5, 9 }

  log.Println("Original values addresses:")
  for _, index := range indexes {
    my_array[index] = "Good element"
    log.Println(&my_array[index], "=", my_array[index])
  }

  log.Println("Selected values addresses:")
  for _, value := range returnSelected(my_array) {
    log.Println(&value, "=", value)
  }

  // Moral of the story is: we don't allocated anything except of array of addresses!
}
