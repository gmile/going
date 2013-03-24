package main

import (
  "log"
)

type MyStruct struct {
  Title string
  Somearray []int
}

func main() {
  var myStruct = new(MyStruct)
  myStruct.Somearray = []int{ 1,2,3 }
  myStruct.Title = "Example"

  log.Println(&myStruct)
  log.Println(&myStruct.Somearray[0])
  log.Println(&myStruct.Somearray[1])
  log.Println(&myStruct.Somearray[2])
  log.Println(&myStruct.Title)

  /*
    Moral of the story:
    - you can get the address of the variable even if it's struct
    - you can't get the address of the array if it's a struct's field
    - you can get the address of the string  if it's a struct's field

    Bonus:
    - even though Title is defined before Somearray in MyStruct,
      it's then implicitly moved to the end of the struct
      (as seen by its address)
  */

  // B
}
