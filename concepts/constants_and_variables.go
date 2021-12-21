package concepts

import "fmt"

// Variables and constants are used to store data.
// As its name suggests, variables can change their values over time, while constants can't.
//
// Golang is a static and strongly typed programming language. It means that all variables have one type that cannot change during runtime.
//
// In Golang, we have the following basic types:
// - bool. True or false values.
// - string. Holds any text.
// - int, int8, int16, int32, int64. Numbers that can be written without a fractional component. These numbers are assigned positive or negative.
// - uint, uint8, uint16, uint32, uint64, uintptr. Numbers that can be written without a fractional component. These numbers aren't assigned positive or negative.
// - byte. It is an alias for uint8
// - rune. It is an alias for int32 and represents a Unicode code point.
// - float32, float64. Numbers that can be written with a fractional component. These numbers are assigned positive or negative.
// - complex64, complex128. Complex numbers.
//
// // To declare constants, you must use the following syntax:
// - const CONSTANT_NAME = VALUE
// where:
// - CONSTANT_NAME is the name of the constant
// - VALUE is the value that should be stored.
//
// To declare variables, you can use one of the following syntaxes:
// - var VARIABLE_NAME TYPE
// - var VARIABLE_NAME = VALUE
// - VARIABLE_NAME := value
// where:
// - VARIABLE_NAME is the name of the variable
// - VALUE is the value that should be stored.
// - TYPE is one of the basic types or one structs (Please check structs.go).

const hello = "Hello, World." // Defining a constant

func ConstantsAndVariables() {
	var number int                  // Declaring an integer variable
	number = 10                     // Assign 10 to this variable
	fmt.Println("Number =", number) // Printing number to the screen
	// number = "maybe?" // if you un-comment, it will raise a compilation error.

	var name = "Matheus"            // Declaring a string variable and assigning a text to it.
	fmt.Println(hello, "I'm", name) // Printing the constant hello and the variable name to the screen
	// hello = "hm..." // if you un-comment, it will raise a compilation error.
	name = "Mina"
	fmt.Println(hello, "I'm", name) // Printing the constant hello and the variable name to the screen

	pi := 3.1415            // Declaring a variable called pi and assigning a value to it
	fmt.Println("Pi =", pi) // Printing pi to the screen
}
