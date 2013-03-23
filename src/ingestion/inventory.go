package ingestion

import (
  "os"
  csv "encoding/csv"
)

type Inventory struct {
  Cars []Car
  // owner?
}

func (i *Inventory) Initialize(csv_path string) {
  file, err := os.Open(csv_path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  matrix, err := csv.NewReader(file).ReadAll()

  i.Cars = make( []Car, len(matrix)-1 )

  for index, row := range matrix[1:] {
    i.Cars[index].Initialize(row)
  }
}

// no other car of the same make
func (i *Inventory) single_cars_cars (cars []Car) {
}

// at least two of the same make exist
func (i *Inventory) make_ad_cars (cars []Car) {
}

// multiple of the same make exist but not model
func (i *Inventory) make_ymm_cars (cars []Car) {
}

// multiple of the same make and model exist
func (i *Inventory) make_model_cars (cars []Car) {
}
