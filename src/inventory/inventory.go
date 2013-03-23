package inventory

import (
  "car"
  "os"
  csv "encoding/csv"
)

type Inventory struct {
  Cars []car.Car
  // owner?
}

func (i *Inventory) Initialize(csv_path string) {
  file, err := os.Open(csv_path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  matrix, err := csv.NewReader(file).ReadAll()

  i.Cars = make( []car.Car, len(matrix)-1 )

  for index, row := range matrix[1:] {
    i.Cars[index].Initialize(row)
  }
}
