package examples

import "fmt"

func Recursion() {
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Printf("Fibbonaci value is: %v\n", fib(7))
}
