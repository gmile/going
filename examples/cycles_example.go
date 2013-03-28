package main

import "log"
import r "math/rand"

func main() {
  type Car struct {
    maker string
    model string
  }

  cars   := [15]Car{}
  makers := [5]string{ "Jeep", "Mercedes", "Porshe", "McLaren", "Audi" }
  models := [5]string{ "Model A", "Model B", "Model C", "Model D", "Model E" }

  for i, _ := range cars {
    cars[i].maker = makers[ r.Intn(len(makers)) ]
    cars[i].model = models[ r.Intn(len(models)) ]
  }

  makes := map[string]map[string][]int {}

  for i, car := range cars {
    car_maker, ok := makes[car.maker]

    if ok {
      _, ok := car_maker[car.model]

      if ok {
        car_maker[car.model] = append(car_maker[car.model], i)
      } else {
        car_maker[car.model] = make([]int, 1)
        car_maker[car.model][0] = i
      }
    } else {
      makes[car.maker] = map[string][]int {}
      makes[car.maker][car.model] = make([]int, 1)
      makes[car.maker][car.model][0] = i
    }
  }

  for maker_title, models := range makes {
    log.Println(maker_title)

    for model, indexes := range models {
      log.Println(" ", model, ":", indexes)
    }
  }
}
