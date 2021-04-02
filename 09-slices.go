package main

import "fmt"

func main()  {
  // slices typed by the elements they contain
  // make built-in creates a slice
  // "zero-valued" for strings is space
  s := make([]string, 3)
  fmt.Println("emp:", s)

  // getting / setting works like arrays
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s) // [a b c]
  fmt.Println("get:", s[2]) // c
  fmt.Println("len:", len(s)) // 3

  // additional built-in support
  // append returns a new slice
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("appended:", s) // [a b c d e f]

  // make new one with same length
  c := make([]string, len(s))
  copy(c, s) // into c from s
  fmt.Println("copy:", c) // [a b c d e f]

  // slice operator takes subseq (range from (inclusive) to (exclusive))
  l := s[2:5]
  fmt.Println("sl1:", l) // [c d e]
  l = s[:5]
  fmt.Println("sl2:", l) // [a b c d e]
  l = s[2:]
  fmt.Println("sl3:", l) // [c d e f]

  // declar and init; notice no size
  t := []string{"g", "h", "i"}
  fmt.Println("t:", t) // [g h i]

  // multi-dimensional slices
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)
  // [
  //  [0]
  //  [1 2]
  //  [2 3 4]
  // ]
}
