package main

import (
  "fmt"
  "math"
)

// declare a constant value
const s string = "constant"

func main()  {
  fmt.Println(s)

  const n = 500000000

  // arbitrary precision
  const d = 3e20 / n
  fmt.Println(d)

  // numberic constants have no type until it's specified
  fmt.Println(int64(d))

  // a number can be given a type by using it in context that requires one
  // math.Sin expects a float64
  fmt.Println(math.Sin(n))
}
