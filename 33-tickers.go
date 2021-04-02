// tickers do something repeatedly at regular intervals
package main

import (
  "fmt"
  "time"
)

func main() {
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)

  go func() {
    // while true
    for {
      select {
        // done has a value?
      case <- done:
        return
      // wait for ticker to send value, and read from it
      // sends time as the value at the interval (500s)
      case t := <-ticker.C:
        fmt.Println("Tick at", t)
      }
    }
  }()

  // give time for goroutine to run 3 times
  time.Sleep(1600 * time.Millisecond)
  ticker.Stop()
  done <- true
  fmt.Println("Ticker stopped")
}
