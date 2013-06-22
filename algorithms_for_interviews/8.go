package main

func ball_breaks_on(floor_number int) bool {
  return floor_number >= 17 // target floor
}

func main() {
  balls         := 0
  current_floor := 0
  lowest_floor  := 0
  highest_floor := 25 // total floors

  for (highest_floor - lowest_floor > 1) {
    current_floor = (highest_floor + lowest_floor) / 2

    balls += 1

    if ball_breaks_on(current_floor) {
      highest_floor = current_floor
    } else {
      lowest_floor = current_floor
    }
  }

  println(balls)
}
