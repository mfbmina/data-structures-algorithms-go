package main

import "fmt"

// In computer science, a loop (or for-loop) is a control flow statement for specifying iteration,
// which allows code to be executed repeatedly.
//
// In Golang, there are three main types of for loop:
// 1. While loops (which sometimes can be infinite loops)
// 2. For loops
// 3. Range loops
func main() {
	array := []int{1, 2, 3, 4, 5}

	// While loops are used to do some actions while the condition is true.
	// In this case, the condition is `i` being below 5.
	i := 0
	for i < 5 {
		fmt.Println(i, "-", array[i])
		i++
	}
	fmt.Println("----------")

	// For loops are used to do some actions when you know exactly all the conditions for it.
	// In this case, we want to instanciate `i` with 0 and for each iteration we increase `i` by 1.
	// We do this while `i` is below 5.
	for i := 0; i < 5; i++ {
		fmt.Println(i, "-", array[i])
	}
	fmt.Println("----------")

	// Range loops are the best when you want to go over a collection of elements.
	// In this case, we go through the array.
	for i, v := range array {
		fmt.Println(i, "-", v)
	}
	fmt.Println("----------")

	// Loops can also be `infinite`
	i = -1
	for {
		i += 1

		if i >= 5 {
			// `break` quits the loop execution
			break
		}

		if i%2 == 0 {
			// `continue` executes the next iteration
			continue
		}

		fmt.Println(i) //, "-", array[i])
	}
}
