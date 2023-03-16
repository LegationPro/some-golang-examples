package examples

import "fmt"

func change_var(some_var *string) {
	*some_var = "not hi anymore >:)"
}

func Pointers() {
	var some_var string = "hi"

	fmt.Printf("Before changing memory: %v\n", some_var)
	change_var(&some_var)
	fmt.Printf("After changing memory: %v\n", some_var)
}
