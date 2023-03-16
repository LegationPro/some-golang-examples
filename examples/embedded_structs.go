package examples

import "fmt"

type Base struct {
	num int
}

func (b Base) describe() string {
	return fmt.Sprintf("base with num=%v\n", b.num)
}

type Container struct {
	base Base
	str  string
}

func EmbeddedStructs() {
	co := Container{
		base: Base{
			num: 1,
		},
		str: "some name",
	}

	type describer interface {
		describe() string
	}

	var d describer = co.base

	fmt.Printf("describe: %v\n", d.describe())
}
