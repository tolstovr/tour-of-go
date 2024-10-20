package main

import (
	"fmt"
)

func main() {
	if age := 17; age >= 18 {
		fmt.Println("You are old enough to drive")
	} else {
		fmt.Printf("You are %d years old, which is not old enough to drive\n", age)
	}
}
