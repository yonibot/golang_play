package main

import (
  "fmt"
  "reflect"
)

var i, j int = 1, 2

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(reflect.TypeOf(i), j, c, reflect.TypeOf(python), reflect.TypeOf(java))
}
