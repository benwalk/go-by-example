package main

import "fmt"

type rect struct {
  width, height int
}

// methods are defined on struct types
// the receiver type is specified before the name of the function
func (r *rect) area() int {
  return r.width * r.height
}
// receiver type could be pointer or value types
func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

func main() {
  r := rect{width: 10, height: 5}
  fmt.Println(r) // {10 5}

  // go automatically handles conversion between values and pointers for method calls
  // pointer reciever types avoid copying and allow the method to mutate the recieving struct
  fmt.Println("area:", r.area()) // 50
  fmt.Println("area:", (&r).area()) // 50
  fmt.Println("perim:", r.perim()) // 30
}
