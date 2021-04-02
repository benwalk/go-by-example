// ranges iterates over elements in some data structures
package main

import "fmt"

func main()  {
  nums := []int{2, 3, 4}
  sum := 0

  // range on arrays and slices provides both the index and the value
  // in this case we discard the index
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum) // 9

  kvs := map[string]string{"a":"apple", "b":"banana"}

  // range on maps returns the k/v pairs
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v) // notice printf usage
  }
  for k := range kvs {
    fmt.Println("key:", k) // a
  }

  // range on strings provides index and unicode
  for i, c := range "go" {
    fmt.Println(i, c)
    // 0 103
    // 1 111
  }
}
