package main

import (
  "math/rand"
  "fmt"
  "time"
)

func quicksort(inputArray []int) (outputArray []int) {
  if len(inputArray) <= 1 {
    return inputArray
  } else {
    pivot_index := len(inputArray) / 2
    pivot := inputArray[pivot_index]
    left  := make([]int, 0)
    right := make([]int, 0)

    for index, e := range inputArray {
      if index != pivot_index {
        if e < pivot {
          left = append(left, e)
        } else {
          right = append(right, e)
        }
      }
    }

    left = append(quicksort(left), pivot)
    return append(left, quicksort(right)...)
  }

  return
}

func main() {
  rand.Seed(time.Now().Unix())
  array := rand.Perm(50)

  fmt.Println(array)
  fmt.Println(quicksort(array))
}
