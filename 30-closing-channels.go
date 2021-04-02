// closing indicates no more values will be sent on a channel
package main

import "fmt"

func main()  {
  jobs := make(chan int, 5)
  done := make(chan bool)

  // worker goroutine reads from jobs while jobs is open,
  // or no more jobs are on the channel
  go func() {
    for {
      // special 2 value form of receive:
      // 'more' will be false if 'jobs' has been closed AND
      // no more values exist in the channel
      j, more := <-jobs
      if more {
        fmt.Println("received job", j)
      } else {
        fmt.Println("received all jobs")
        // notify done on separate channel
        done <- true
        return
      }
    }
  }() // kick off worker

  // queue jobs
  for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("send job", j)
  }
  // built-in
  close(jobs)

  // block until done receives a value (and then read it)
  <-done
}
