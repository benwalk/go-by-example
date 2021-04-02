package main

import "fmt"

func main()  {

  // basic: single condition
  i := 3
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // without a condition, it's a while
  // need a break or return
  for {
    fmt.Println("loop")
    break
  }

  // contine can be used to jump to next iteration
  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n) // prints odds
  }
}
