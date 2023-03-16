package examples

import "fmt"

func return_multiple_vals() (int32, string, string) {
	return 2, "hi", "i'm good"
}

func MultipleReturn() {
	n0, n1, n2 := return_multiple_vals()

	fmt.Printf("Multiple returns: \n0: %v\n 1: %v\n 2: %v\n", n0, n1, n2)
}
