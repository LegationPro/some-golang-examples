package examples

import "fmt"

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func measure(g geometry) int {
	return int(g.area())
}

func Interfaces() {
	r := rect{
		width:  150,
		height: 250,
	}

	area := r.area()

	fmt.Printf("Area is: %v\n", area)

	// Must be passed by ref to convert
	result := measure(&r)

	fmt.Printf("Result is: %v\n", result)
}
