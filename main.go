package main

import "fmt"

func main() {
	sum := 0

	for {
		sum += 1
		fmt.Println(sum)

		if sum >= 50 {
			fmt.Println("OKAY!")
			break
		}
	}
}
