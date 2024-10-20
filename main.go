package main

import "fmt"

func add(a, b int) int { // a and b has same type, so we can shorten our code
	return a + b
}

func printNum(sum int, prefix string) {
	fmt.Println(prefix, sum)
}

func main() {
	printNum(add(40, 2), "Answer is:")
}
