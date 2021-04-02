// idiomatic to communicate erros via an explicit, separate return value
// unlike java, ruby, and overloaded single result object used elsewhere

package main

import (
  "errors"
  "fmt"
)

// by convention, errors are the final return value and have type 'error'
func f1(arg int) (int, error) {
  if arg == 42 {
    // errors.New constructs a basic error value with a given message
    return -1, errors.New("can't work with 42")
  }

  // a nil value in the error position indicates that there was no error
  return arg + 3, nil
}

// custom error types can be used if the Error() method is defined on them
type argError struct {
  arg int
  prob string
}
// defines Error() by taking argError as a reciever
func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
  if arg == 42 {
    // constructs the error and produces a pointer to it
    return -1, &argError{arg, "can't work with it"}
  }
  return arg + 3, nil
}

func main()  {
  // index is discarded, the value is i
  for _, i := range []int{7, 42} {
    // idiomatic: calls function and immediately check if error exists
    if r, e := f1(i); e != nil {
      fmt.Println("f1 failed:", e) // when i == 42
    } else {
      fmt.Println("f1 worked:", r) // when i == 7
    }
  }

  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed:", e) // when i == 42
    } else {
      fmt.Println("f2 worked:", r) // when i == 7
    }
  }

  // to use data in a custom error, need to get an instance of the error from the reference
  _, e := f2(42)
  // type assert error to `argError` and if successful
  if ae, ok := e.(*argError); ok {
      fmt.Println(ae.arg) // 42
      fmt.Println(ae.prob) // can't work with it
  }
}
