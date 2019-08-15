package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its weekday")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Its boolean.")
		case int:
			fmt.Println("Its Integer.")
		default:
			fmt.Println("Its neither boolean nor Integer.")
		}
	}
	whatAmI(true)
	whatAmI(10)
	whatAmI(100.89)
	whatAmI("String")
}
