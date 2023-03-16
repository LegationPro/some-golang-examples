package examples

import "fmt"

type SomeStruct struct {
	name      string
	last_name string
}

func (s *SomeStruct) get_name() string {
	return s.name
}

func (s *SomeStruct) get_lastname() string {
	return s.last_name
}

func Structs() {
	some_struct := SomeStruct{
		name:      "Anze",
		last_name: "Johnson",
	}

	lastname := some_struct.get_lastname()
	name := some_struct.get_name()

	fmt.Printf("My name is: %v\n", name)
	fmt.Printf("My last name is %v\n", lastname)
}
