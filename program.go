package main

import (
  "fmt"
  "os"
  convert "strconv"
  csv "encoding/csv"
  m   "my_structs"
)

func readCsv(path string) []m.Car {
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  matrix, err := csv.NewReader(file).ReadAll()

  cars := make( []m.Car, len(matrix)-1 )

  for i, row := range matrix[1:] {
    cars[i].Model = row[6]
    cars[i].Year, err = convert.ParseUint(row[4], 10, 0)
  }

  return cars
}

func main() {
  var cars []m.Car = readCsv("/Users/gmile/odesk/haystack_profiling/homenet.txt")

  for _, car := range cars[1:10] {
    fmt.Println(car.Model, "was made in", car.Year)
  }
}
