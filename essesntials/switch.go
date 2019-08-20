package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend Boooy!")
	default:
		fmt.Println("Week day!cwork work :/")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its Morning")
	default:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch fu := i.(type) {
		case bool:
			fmt.Println("I'm boolean!")
		case int:
			fmt.Println("I'm integer!")
		default:
			fmt.Println("Don't know type  %T\n", fu)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("yes")
}
