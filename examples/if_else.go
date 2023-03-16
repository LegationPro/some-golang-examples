package examples

import "fmt"

func IfElse() {
	var i uint64 = 0
	var j uint64 = 4

	if i == 0 {
		fmt.Printf("\ni is %v\n", i)
	}

	if j == 3 {
		fmt.Printf("\noh its 3 = %v\n", j)
	} else if j == 4 {
		fmt.Printf("\nohhh its 4 = %v\n", j)
	} else {
		fmt.Printf("\n oh it's not 3 or 4 its: %v\n", j)
	}
}
