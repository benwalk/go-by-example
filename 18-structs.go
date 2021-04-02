package main

import "fmt"

// structs are typed collections of fields
// useful for grouping data to form records
type person struct {
  name string
  age int
}

func newPerson(name string) *person {
  // partial construction
  p := person{name: name}

  // setting field value
  p.age = 42

  // local var data still exists and thus a pointer is safe to return
  return &p
}

func main()  {
  // basic construction: fields inferred by order
  fmt.Println(person{"Bob", 20}) // {Bob 20}

  // named field construction:
  fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

  // omitted fields are zero-valued
  fred := person{name: "Fred"}
  fmt.Println(fred.age) // 0

  // & prefix yields a pointer to the struct
  fmt.Println(&person{"Ann", 40}) // &{Ann 40}

  // using a constructor is idiomatic
  fmt.Println(newPerson("Jon")) // &{Jon 42}

  // field access using dots also works with pointers - automatic dereference
  seanPtr := newPerson("Sean")
  fmt.Println(seanPtr.name) // Sean

  // structs are mutable
  seanPtr.age = 51
  fmt.Println(seanPtr) // &{Sean 51}
}
