package main

import (
  "encoding/json"
  "fmt"
  // "os"
)

type response1 struct {
  Page    int
  Fruits  []string
}

type response2 struct {
  Page    int       `json:"page"`
  Fruits  []string  `json:"fruits"`
}

func main() {
  bolB, _ := json.Marshal(true)
  fmt.Println(string(bolB)) // true

  intB, _ := json.Marshal(1)
  fmt.Println(string(intB)) // 1

  strB, _ := json.Marshal("gopher")
  fmt.Println(string(strB)) // "gopher"

  // slices and maps encode to json arrays and objects
  slcD := []string{"apple", "peach", "pear"}
  slcB, _ := json.Marshal(slcD)
  fmt.Println(string(slcB))
}
