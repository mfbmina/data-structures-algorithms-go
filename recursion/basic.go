package recursion

import "fmt"

// In short, recursion is when a function call itself.
//
// When using recursion is important to always have a stop condition
// because if there isn't one, it will consume all your memory.
func Basic(n int64) {
	fmt.Println("Calculating fatorial for", n)
	resp := fat(n)
	fmt.Println("The response is", resp)
	fmt.Println("Calculating fibonacci for", n)
	resp = fib(n)
	fmt.Println("The response is", resp)
}

func fat(n int64) int64 {
	// stop condition
	if n == 1 {
		return 1
	}

	// calling itself again!
	return n * fat(n-1)
}

func fib(n int64) int64 {
	// stop condition
	if n > 1 {
		// calling itself again!
		return fib(n-1) + fib(n-2)
	}

	return 1
}
