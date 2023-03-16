package examples

import "fmt"

func add(a, b int) int {
	return a + b
}

func Functions() {
	fmt.Printf("2 + 3 is: %v\n", add(2, 3))
}
