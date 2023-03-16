package examples

import "fmt"

// K(key) comparable - means it is comparable (!= ==)
// V(value) any - alias for interface{}
// []K - we return a array of type of K

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// T - represents a generic
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elements []T

	for e := list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func Generics() {
	var m = map[int]string{
		1: "2",
		2: "2",
		3: "4",
	}

	fmt.Printf("keys:%v", MapKeys(m))

	_ = MapKeys(m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
