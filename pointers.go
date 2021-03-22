package main

import "fmt"

// From the Golang documentation:
// Go has pointers. A pointer holds the memory address of a value.
//
// The type *T is a pointer to a T value. Its zero value is nil.
// Unlike C, Go has no pointer arithmetic.

type Rectangle struct {
  lenght int
  width int
}

func main() {
  // Defining a pointer p.
  var p *int
  i := 42

  // If you want to get the pointer for another variable you can use the & operator.
  p = &i

  // To access value from pointer you should use the * operator.
  fmt.Println("Value holded by p", p)
  fmt.Println("Reading i through p", *p)

  // setting i through the pointer p
  // This is known as "dereferencing" or "indirecting".
  *p = 21
  fmt.Println("New value at i through p", *p)

  r := Rectangle{1, 2}
  // Defining a pointer for r
  r_pointer := &r

  // If you have a pointer for a struct, you can call its attributes in the same way if it was not a pointer.
  area := r_pointer.lenght * r_pointer.width
  fmt.Println("rectangle area's:", area)
}
