package examples

import "fmt"

func sum(nums ...int) {
	sumup := 0

	for n := range nums {
		sumup += nums[n]
	}

	fmt.Printf("The sum of all numbers is: %v\n", sumup)
}

func VariadicFunctions() {
	sum(1, 2, 3, 4, 5, 6, 7)
	nums := []int{1, 2, 35, 6}
	sum(nums...)
}
