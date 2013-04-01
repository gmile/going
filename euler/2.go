package main

func main() {
  x := 2
  fibs := [3]int{ 1, 2, 0 }

  for fibs[0] + fibs[1] < 4000000 {
    fibs[2] = fibs[0] + fibs[1]
    fibs[0], fibs[1] = fibs[1], fibs[2]

    if fibs[2] % 2 == 0 {
      x += fibs[2]
    }
  }

  println(x)
}
