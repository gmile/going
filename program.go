package main

import (
  "fmt"
  "os"
  csv "encoding/csv"
  "car"
)

func readCsv(path string) []car.Car {
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  matrix, err := csv.NewReader(file).ReadAll()

  cars := make( []car.Car, len(matrix)-1 )

  for i, row := range matrix[1:] {
    cars[i].Initialize(row)
  }

  return cars
}

func main() {
  var cars []car.Car = readCsv("/Users/gmile/odesk/haystack_profiling/homenet.txt")

  for _, car := range cars[1:5] {
    fmt.Println(car.Model, "was made in", car.Year)
  }
}
