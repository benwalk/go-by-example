// select lets you wait on multiple channel operations
package main

import (
  "fmt"
  "time"
)

func main() {

  c1 := make(chan string)
  c2 := make(chan string)

  // anon function that waits 1s and then sends to c1
  go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
  }() // immediate exec

  // anon function that waits 2s and then sends to c2
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
  }() // immediate exec

  // each channel receives a value at different times

  for i := 0; i < 2; i++ {
    fmt.Println(i)

    // await on all cases until received
    select {
    // this one comes first, then loop continues
    case msg1 := <-c1:
      fmt.Println("received", msg1)
    // this one comes on second interation
    case msg2 := <-c2:
      fmt.Println("received", msg2)
    }
  }
}
