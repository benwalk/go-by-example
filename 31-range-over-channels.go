package main

import "fmt"

func main()  {
  // open chan with 2 buffer
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)

  // range iterates over values in a channel
  for elem := range queue {
    fmt.Println(elem)
  }
}
