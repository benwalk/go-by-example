package main

import (
  "fmt"
  "time"
)

// takes a channel of booleans for the purpose of indicating work is done
func worker(done chan bool) {
  fmt.Println("working...")
  time.Sleep(time.Second*1)
  fmt.Println("done")

  done <- true
}

func main() {
  // make a channel of booleans for tracking if work is done, buffering 1 work unit
  done := make(chan bool, 1)

  // kick off a worker, which ends up blocked until a reciever is created on the channel
  go worker(done)

  // receive from done
  fmt.Println("waiting for worker")
  <- done
  fmt.Println("worker done")
}
