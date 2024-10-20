package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // same as switch true. Can substitute long if-else chain
	case t.Hour() <= 12:
		fmt.Println("Good morning!")
	case t.Hour() <= 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
