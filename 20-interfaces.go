// interfaces are named collections of method signatures
package main

import (
  "fmt"
  "math"
)

// basic interface for geometric shapes
type geometry interface {
  area() float64
  perim() float64
}

// some geometric shapes
type rect struct {
  width, height float64
}
type circle struct {
  radius float64
}

// to implement an interface, implement all methods defined in the interface
func (r rect) area() float64  {
  return r.width * r.height
}
func (r rect) perim() float64 {
  return r.width*2 + r.height*2
}
func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
  return math.Pi * c.radius * 2
}

// if a variable (g) has an interface type
// methods in that interface can be executed
// variables that don't have all methods of an interface defined cannot
// be passed as arguments to functions that take interfaces
func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

func main()  {
  r := rect{3, 4}
  measure(r) // g: {3 4}, area: 12, perim: 14

  c := circle{radius: 5}
  measure(c) // g: {5}, area: 78.53981633974483, perim: 31.41592653589793
}
