package main

import "fmt"

func main() {

  // creates an array of given size (5) holding ints
  // default values are "zero-valued" which means 0 for ints
  var a [5]int
  fmt.Println("emp:", a) // [0 0 0 0 0]

  // setting and accessing use [] syntax, built-in for length
  a[4] = 100
  fmt.Println("set:", a) // [0 0 0 0 100]
  fmt.Println("get:", a[4]) // 100
  fmt.Println("len:", len(a)) // 5

  // init during decla
  b := [5]int{1,2,3,4,5} // [1 2 3 4 5]

  // multi-dimensional arrays like java
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD) // [[0 1 2] [1 2 3]]
}
