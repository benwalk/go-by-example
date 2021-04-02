package main

import (
  "fmt"
  "time"
)

func main()  {
  // timers represent an event in the future
  // creates a channel that will receive a value
  // in the specified time
  timer1 := time.NewTimer(time.Second * 2)

  // blocks until the timer (channel) receives a value
  <-timer1.C
  fmt.Println("Timer 1 fired")

  // timers are an alternative to sleep, but with the benefit
  // of being stoppable
  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("Timer 2 fired")
  }() // starts a 1s timer

  // stop immediately (before it fires)
  stop2 := timer2.Stop()

  if stop2 {
    fmt.Println("Timer 2 stopped")
  }

  time.Sleep(time.Second * 2)
}
