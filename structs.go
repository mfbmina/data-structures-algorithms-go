package main

import "fmt"

// Structs are an abstraction for a collection of fields and it is mostly used for defining new types.
//
// Defining a struct in Golang is simple. The syntax is the following:
// type STRUCT_NAME struct {
//   FIELD_NAME FIELD_TYPE
// }
// where:
// - STRUCT_NAME: the name that you're gonna use for the struct.
// - FIELD_NAME: the name that you're gonna use for the field.
// - FIELD_TYPE: any of default types in Go or structs that you've created.
type Rectangle struct {
  lenght int
  width int
}

func main() {
  var rectangle_1 Rectangle // Declaring a struct without any default value.
  fmt.Println(rectangle_1)

  // To set the fields of the struct, you use `variable.field_name` and then the operator =
  rectangle_1.width = 1
  rectangle_1.lenght = 2
  fmt.Println(rectangle_1)

  rectangle_2 := Rectangle{3, 2} // Declaring a struct with default values.

  // To get the fields of the struct, you use `variable.field_name`
  area := rectangle_2.lenght * rectangle_2.width

  fmt.Println("rectangle_2 area's:", area)
}
