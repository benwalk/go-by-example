package main

import "fmt"

// takes two ints and returns sum
func plus(a int, b int) int {
  return a + b
}

// conseq same-typed params can omit type
func plusPlus(a, b, c int) int {
  return a + b + c
}

func main()  {
  res := plus(1,2)
  fmt.Println("1 + 2 = ", res) // 3

  res = plusPlus(1, 2, 3)
  fmt.Println("1 + 2 + 3 = ", res) // 6
}
