// basic send and receives on channels are blocking
// select with a default enables nonblocking flows
package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  // takes the first
  select {
    // if a msg is ready and available
  case msg := <-messages:
    fmt.Println("received message", msg)
    // otherwise no message is ready
  default: fmt.Println("no message received")
  }
  // prints "no message received" because no msg ready

  msg := "hi"
  select {
  // cannot send because messages is unbuffered
  case messages <- msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }
  // prints "no message sent" because cannot send

  select {
  // read from messages?
  case msg := <-messages:
    fmt.Println("received message", msg)
  // read from signals?
  case sig := <-signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }
  // prints "no activity" because no messages or signals sent
}
