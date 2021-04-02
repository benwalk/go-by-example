package main

import "fmt"

func main() {
  // var declares variables
  var a = "initial"
  fmt.Println(a)

  // multiple
  var b, c int = 1, 2
  fmt.Println(b, c)

  // types inferred
  var d = true
  fmt.Println(d)

  // declared without init are "zero-valued" aka default
  var e int
  fmt.Println(e) // 0

  // := declare and init at the same time
  var f string = "apple"
  g := "orange"
  fmt.Println(f)
  fmt.Println(g)
}
