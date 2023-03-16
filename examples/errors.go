package examples

import (
	"errors"
	"fmt"
)

func isNumber1OrError(n int) (error, bool) {
	if n == 1 {
		return nil, true
	}

	err := errors.New("i is not 1")

	return err, false
}

func Errors() {
	num := 1

	err, state := isNumber1OrError(num)
	if err != nil {
		panic(err)
	}

	if err == nil {
		fmt.Printf("No errors! state is: %v\n", state)
	}

	err2, state2 := isNumber1OrError(2)
	if err2 != nil {
		panic(err2)
	}

	if err2 == nil {
		fmt.Printf("No errors! state is: %v\n", state2)
	}
}
