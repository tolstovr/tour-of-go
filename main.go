package main

import "fmt"

func add(a, b int) int { // a and b has same type, so we can shorten our code
	return a + b
}

func swap(x, y int) (int, int) { // we return two values here, which can be assigned to multiple variables
	return y, x
}

func main() {
	a, b := 5, 10
	fmt.Println("Before swapping:", a, b)
	a, b = swap(a, b)
	fmt.Println("After swapping:", a, b)
}
