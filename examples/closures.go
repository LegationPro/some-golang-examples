package examples

import "fmt"

func closure() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Closures() {
	i_is := closure()
	fmt.Printf("i is: %v\n", i_is())
}
