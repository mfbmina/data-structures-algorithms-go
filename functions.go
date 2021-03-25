package main

import "fmt"

// Functions can be understand as a group of instructions.
// Using functions is a good practice while programming because you can transform one huge set of instructions
// into smaller ones. The benefits includes:
// - readability
// - testability
// - maintanability
// - reusability
//
// To declare an array in Golang, you can use the following syntaxes:
// func FUNCTION_NAME(PARAMETERS) RETURN_TYPE { CODE }
// where:
// - FUNCTION_NAME: the name that you're gonna use for the function.
// - PARAMETERS: optional comma separated list of parameters. Eg. number int, int year
// - RETURN_TYPE: optional. It defines the type of the return of the funcion.
// - CODE: any group of instructions
//
// To run a Go code, we always define a function called main, which will be executed when you call:
// $ go run file.go
func main() {
   // This is how a function can be called.
  hi()

  // The return of functions can be stored as variables.
  // Also, to provide the arguments to the functions you don't need to say the type.
  value := add(1, 1)
  fmt.Println("1 + 1 =", value)

  x := 1
  y := 2
  // This is how we call a function with pointers.
  // We can see the variables from main function got their value changed.
  swap(&x, &y)
  fmt.Println("x =", x)
  fmt.Println("y =", y)

  // Passing array to a function.
  array := []int{1, 2, 3, 4, 5}
  print_elements(array, 5)
}

// We are declaring a function called `hi`, which will add two numbers.
// As you can see, it receives no parameters and will not return anything to the main function.
func hi() {
  fmt.Println("Hello!")
}

// We are declaring a function called `add`, which will add two numbers.
// As you can see, it receives two parameters of `int` type: `x` and `y`.
// It will always returns another `int`.
func add(x int, y int) int {
  return x + y
}

// Functions can receive pointers as arguments.
// This is really useful when we want to change values in variables from different contexts.
// In this case, we are receive pointers for two integers and we are swaping their values.
func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

// Functions can receive arrays as arguments.
func print_elements(array []int, size int) {
  for i := 0; i < size; i++ {
    fmt.Println(array[i])
  }
}
