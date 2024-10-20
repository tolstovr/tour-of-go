package main

import (
	"fmt"
)

func main() {
	defer fmt.Print("World!\n") // defers the execution of the following statement until the surrounding function returns
	fmt.Print("Hello, ")
}
