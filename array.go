package main

import "fmt"

// Arrays can be considered a list of elements with a defined size.
// It is not possible to change the size of an array later.
//
// To declare an array in Golang, you should the syntax is the following:
// var VAR_NAME [SIZE]TYPE{DEFAULT_VALUES}
// where:
// - VAR_NAME: the name that you're gonna use for the variable.
// - SIZE: how many elements that array will store.
// - TYPE: any of default types in Go or structs that you've created.
// - DEFAULT_VALUES: this is optional but it is used when you want to create an array with default values on it.
func main() {
  var words [2]string // declaring an array called words with 2 elements.
  words[0] = "Hello" // assign the word "Hello" to the first element into the array
  words[1] = "World" // assign the word "World" to the second element into the array

  fmt.Println(words[0], words[1]) // Prints the first and the second element of the array.
  fmt.Println(words) // Prints the whole array

  primes := [6]int{2, 3, 5, 7, 11, 13} // declaring an array for the 6 prime numbers.
  fmt.Println(primes)
}
