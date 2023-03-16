package examples

import "fmt"

func Arrays() {
	var arr [5]int
	fmt.Printf("first element of array is: %v\n", arr[0])

	b := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	for _, v := range b {
		fmt.Printf("current element of array is: %v\n", v)
	}
}
