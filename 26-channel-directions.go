// restricting chan directions helps add type safety
package main

import "fmt"

// only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
  pings <- msg
}

// accepts a channel to read from and a channel to send to
func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings // reads from pings
  pongs <- msg // sends it to pongs
}

func main()  {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message") // sends to pings
  pong(pings, pongs) // transfers from pings to pongs
  fmt.Println(<-pongs) // reads from pongs
}
