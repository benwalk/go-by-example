package main

import "fmt"

func main()  {
  // if / else
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  // if, no else
  if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  // statement can precede conditionals
  // vars available in all branches
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits.")
  }

  // no parens req but braces are
  // there's no ternary operator, so a full if else is required

}
