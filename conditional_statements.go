package main

import "fmt"

// Conditionals help us to choose an action based on the data that we have.
//
// To have conditional statements, we need to have our conditional operators.
// Conditional operators tell you which comparisons you are allowed to do.
// In Golang, we have the following operators:
// - `A == B`: returns `true` if A is equal to B
// - `A != B`: returns `true` if A is different from B
// - `A >  B`: returns `true` if A is higher than B
// - `A >= B`: returns `true` if A is higher or equal to B
// - `A <  B`: returns `true` if A is lower than B
// - `A <= B`: returns `true` if A is lower or equal to B
//
// We also have logical operators:
// - `X && Y`: `&&` is the logical operator `AND`. It expects both expressions to be true.
// - `X || Y`: `||` is the logical operator `OR`. It expects one expression to be true.
// - `!X`: `!` is the operator `NOT`. It negates the value of the expression. If the expression is already false, it becomes true. `!false == true`.
//
// In Golang we have two ways of doing conditional statements: `if/else` and `switch/case`.
func main() {
  var age int
  var country string
  fmt.Println("How old are you?")
  fmt.Scanln(&age)
  fmt.Println("Where are you from?")
  fmt.Scanln(&country)

  // One good example is asking your age before visiting some websites.
  // If the person is a minor, we shouldn't allow him/her to visit the page.
  // This case is what we call `if/else` statement because it should choose between actions by evaluating expressions as true or false.
  // It will always execute the first true condition and if no condition is satisfied, it will execute the else.
  if age >= 21 {
    // If the provided age is equal or higher to 21, it will print "Welcome!"
    fmt.Println("Welcome!")
  } else if age >= 18 && country != "US" {
    // If the provided age is equal or higher to 18 and the user is not from the US it will print the message.
    fmt.Println("Quite young but still welcome!")
  } else {
    // If the age is below 18 it will print the following message.
    fmt.Println("hmmm... maybe you should not be allowed on here!")
  }

  var animal string
  fmt.Println("What is your favorite animal?")
  fmt.Scanln(&animal)

  // Sometimes, we do not want to check if something is true or false, but to check if it has a specific value.
  // This is the best option to use the switch statement.
  switch animal {
  case "dog":
    fmt.Println("Dogs bark, right?")
  case "cat":
    fmt.Println("Meow!")
  case "fish":
    fmt.Println("I've no clue how a fish sounds...")
  default:
    fmt.Println("Hm... I need to upgrady my animal list...")
  }
}
