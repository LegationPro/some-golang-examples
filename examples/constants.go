package examples

import "fmt"

func Constants() {
	const this_one_variable = "I'm a constant!"

	const (
		something      string = "Something constant"
		something_else string = "Something else. i'm a constant too"
		this_is_cool   string = "This is so cool! i'm constant"
	)

	fmt.Printf("%v %v %v", something, something_else, this_is_cool)
}
