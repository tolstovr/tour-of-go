package main

import "fmt"

func main() {
	var a [2]string // [size]type is an array

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var slice []int = primes[1:4] // []type is a slice. Same as in Python. Also with defaults [start:], [:end], [:]
	fmt.Println(slice)
	fmt.Println(primes[1:4]) // it's also okay

	primes[2] = 1000
	fmt.Println(slice) // changed!

	fmt.Println(len(slice), cap(slice)) // len -- length, cap -- capacity (how many elements can be added without resizing)

	var s []int
	fmt.Println(s, len(s), cap(s))

	v := make([]int, 5, 5) // creating a slice with length 5 and capacity 5
	fmt.Println(v)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][1] = "O"
	board[0][2] = "X"
	board[2][0] = "X"
	fmt.Println(board)

	v = append(v, 5)
	fmt.Println(v, len(v), cap(v))

	for index, value := range primes {
		fmt.Println(index, value)
	}

	for _, value := range primes {
		fmt.Println(value)
	}

	for i := range primes {
		fmt.Println(i)
	}
}
