package main

import "fmt"

func add(a, b int) int { // a and b has same type, so we can shorten our code
	return a + b
}

func swap(x, y int) (int, int) { // we return two values here, which can be assigned to multiple variables
	return y, x
}

func aboba(num int) (x, y int) { // values of x and y will be returned
	x = num * 2
	y = num * 3
	return // same as return x, y
}

func main() {
	a, b := 5, 10
	fmt.Println("Before swapping:", a, b)
	a, b = swap(a, b)
	fmt.Println("After swapping:", a, b)

	fmt.Println(aboba(50))
}
