package main

import "fmt"

// by default channels are unbuffered meaning
// they only accept sends (chan <-) IF there is a corresponding
// receive (<- chan) ready to receive the sent value

// buffered channels accept a limited number of values without a corresponding
// receiver for those values
func main()  {
  // make a channel accepting strings buffering up to 2 values
  messages := make(chan string, 2)

  // without a corresponding concurrent channel reader, messages can be buffered
  messages <- "buffered"
  messages <- "channel"
  // messages <- "overflow" this causes a deadlock because there are no goroutines to receive and buffer is full

  // receiving those values can happen asynchronously
  fmt.Println(<-messages) // "buffered"
  messages <- "overflow" // can happen now that buffer is 1/2
  fmt.Println(<-messages) // "channel"

  // notice FIFO ordering
}
