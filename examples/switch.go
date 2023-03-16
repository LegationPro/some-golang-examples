package examples

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2

	switch i {
	case 1:
		fmt.Println("\n i is 1!")
	case 2:
		fmt.Println("\n i is 2!")
	default:
		fmt.Printf("\n i is something else! which is: %v \n", i)
	}

	t := time.Now()

	switch {
	case t.Hour() > 12:
		fmt.Printf("it's before noon!\n Hour: %v\n", t.Hour())
	default:
		fmt.Printf("It's after noon!\n Hour: %v\n", t.Hour())
	}

	whatTypeAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("\ni'm a boolean!!")
		case int:
			fmt.Println("\nI'm a number!")
		default:
			fmt.Printf("Don't know what type \n I am: %T\n", t)
		}
	}

	whatTypeAmI(true)
	whatTypeAmI(2)
	whatTypeAmI("hey")
}
