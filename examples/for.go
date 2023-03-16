package examples

import "fmt"

func For() {
	i := 1

	for i <= 3 {
		fmt.Println(uint64(i))
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	loop_start_num := 0

	for {
		fmt.Printf("\n Loop! %v\n", loop_start_num)
		loop_start_num += 1
		if loop_start_num >= 10 {
			break
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Printf("%v\n", n)
	}
}
