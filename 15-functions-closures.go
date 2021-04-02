package main

import "fmt"

// function that returns an unnamed function
// anon function encloses 'i'
func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  // nextInt is a function
  // when called, enclosed i is incremented and returned
  nextInt := intSeq()
  fmt.Println(nextInt()) // 1
  fmt.Println(nextInt()) // 2
  fmt.Println(nextInt()) // 3

  // state of i is unique to that function
  otherSeq := intSeq()
  fmt.Println(otherSeq()) // 1
}
