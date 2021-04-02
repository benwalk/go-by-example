package main

import "fmt"

// param is an int, so args passed are values
// ival is a copy of whatever is passed from calling function context
func zeroval(ival int) {
  ival = 0
}

// * indicates pointer to a type as a func parameter
func zeroptr(iptr *int) {
  // *iptr de-references the memory address
  // assigning a value to a de-referenced pointer updates the value at the address
  *iptr = 0

  // iptr = 0 is wrong
}

func main()  {
  i := 1
  fmt.Println("initial:", i) // 1

  // the value of i (1) is passed
  zeroval(i) // enclosed 'ival' is set to 0
  fmt.Println("zeroval:", i) // prints the i set above (1)

  // & is address of a var, iow, a pointer to i
  zeroptr(&i) // expects a pointer, de-references pointer to update value to 0
  fmt.Println("zeroptr:", i) // 0

  // pointers can be printed too
  fmt.Println("pointer:", &i) // 0xc00001e078

}
