package main

import (
  "fmt"
  "time"
)

// takes a jobs channel *reader* consisting of ints
// takes a results channel *writer* only allowing ints
func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "started  job", j)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished jbo", j)
    results <- j * 2
  }
}

func main() {
  const numJobs = 5
  jobs := make(chan int, numJobs) // jobs channel with buffer 5
  results := make(chan int, numJobs) // results channel with buffer 5

  // fire off 3 workers who will wait for jobs to contain stuff to do
  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  // add 5 jobs to the channel for the workers
  for j := 1; j <= numJobs; j++ {
    jobs <- j
  }
  close(jobs) // closing indicates no more jobs will come

  // since we added 5 jobs, let's check for 5 results
  for a := 1; a <= numJobs; a++ {
    <- results
  }

  // why doesn't worker 1 get all the jobs?
  // worker 2 will take from jobs while worker 1 is sleeping
  // worker 3 will take from jobs while worker 2 is sleeping
}
