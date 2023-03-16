package examples

import "fmt"

func Slices() {
	// initially set with 0 containing values
	s := make([]string, 3)

	// set accordingly

	s[0] = "hi"
	s[1] = "my name is"
	s[2] = "anze!"

	for _, v := range s {
		fmt.Printf("\ncurrent element of slice is: %v\n", v)
	}
}
