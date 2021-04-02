// "pipes" that connect concurrent goroutines
// send values into channels from one goroutine and receive in another goroutine
package main

import "fmt"

func main()  {

  // channels are typed by the values they convey and made with built-in make
  messages := make(chan string)

  // send a value with anon func using 'channel <-' syntax
  go func() { messages <- "pong" }()
  // alternative:
  go ping(messages)

  // read value with '<-channel'
  msg := <-messages
  fmt.Println(msg)
  msg = <-messages
  fmt.Println(msg)

  // by default, sends and recieves block until both the sender and the reciever
  // are ready
}

// alternative to anon function
func ping(c chan string) {
  c <- "ping"
}
