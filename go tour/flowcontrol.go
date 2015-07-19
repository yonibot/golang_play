package main

import (
  "fmt"
  // "reflect"
)

func main() {
  // sum := 0
  // only looping construct in Go is 'for'
  // for i := 0; i < 10; i++ {
  //   sum += i
  // }

  // leave pre/post statements empty
  // sum := 1
  // for ; sum < 1000; {
  //   sum += sum
  // }
  // -->>> ABOVE is equivalent to:
  // sum := 1
  // for sum < 1000 {
  //   sum += sum
  // }

  // if statement
  // x := 5
  // if x < 10 {
  //   fmt.Println("X is less than 10", reflect.TypeOf(x))
  // }

  // initialize a value -> vars available in all branches of if/else
  if x := 5; x > 4 {
    fmt.Println(x, "is greater than 4.")
  }

  // fmt.Println(x)
}
