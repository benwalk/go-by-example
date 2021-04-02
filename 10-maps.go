// maps are associative data type, a la hashes, dicts, etcs
package main

import "fmt"

func main()  {
  // maps [key-type] to value-type
  m := make(map[string]int)

  // set similar to arrays and slices
  m["k1"] = 7
  m["k2"] = 13

  // prints all k/v pairs
  fmt.Println("map:", m) // map[k1:7 k2:13]

  v1 := m["k1"]
  fmt.Println("v1:", v1) // 7
  fmt.Println("len:", len(m)) // 2

  // removes a pair
  delete(m, "k2")
  fmt.Println("map:", m) // map[k1:7]

  // optional second return value on a get indicates existence
  // blank identifier is a way to discard or ignore it
  _, exists := m["k2"]
  fmt.Println("exists:", exists) // false

  // declar and init
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("nap:", n) // map[bar:2 foo:1]

}
