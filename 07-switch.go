package main

import (
  "fmt"
  "time"
)

func main()  {

  // basic switch
  i := 2
  fmt.Println("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // multiple expressions in the same case
  // keyword default used to catch all
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("It's a weekday.")
  }

  // an alternative to if/else
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon.")
  default:
    fmt.Println("It's after noon.")
  }

  // a type switch compares types instead of values
  // used to discover the type of an interface value
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Println("Not sure what type %T\n", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("string")
}
