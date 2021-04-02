package main

import (
  "fmt"
  "time"
)

func main() {
  // buffered
  c1 := make(chan string, 1)

  go func() {
    time.Sleep(time.Second * 2)
    // c1 is buffered, so this is non blocking
    c1 <- "result 1"
  }()

  // await for whichever comes first
  select {
  // c1 received a message?
  case res := <-c1:
    fmt.Println(res)

  // 1s passed
case <-time.After(time.Second * 1):
    fmt.Println("timeout 1")
  }
  // because the c1 send takes 2s, timeout happens first

  c2 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 1)
    c2 <- "result 2"
  }()

  select {
  // c2 received a message?
  case res := <-c2:
    fmt.Println(res)
  // 2s passed
  case <-time.After(time.Second * 2):
      fmt.Println("timeout 2")
  }
  // c2 gets a message before timeout

}
