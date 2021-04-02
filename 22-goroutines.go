// goroutine is a lightweight thread of execution
package main

import (
  "fmt"
  "time"
)

// prints a "string : n", from 0 -> 2
func f(s string) {
  for i := 0; i < 3; i++ {
    fmt.Println(s, ":", i)
  }
}

func main() {
  // direct call
  f("direct")
  // direct : 0
  // direct : 1
  // direct : 2

  // go keyword executes concurrently with calling function
  go f("goroutine")

  // can also use anon funcs
  go func(msg string) {
    fmt.Println(msg)
  }("going") // immediately call func with arg "going"

  // "going" should be interleaved with "goroutine" if run multiple times.

  // wait and finish
  time.Sleep(time.Second)
  fmt.Println("done")
}
