package main

import "fmt"

func main() {
	v := 42           // type is int
	new := v          // type is int as well
	complex := 2 + 4i // type is complex128

	fmt.Printf("Type of v is %T\n", v)
	fmt.Printf("Type of new is %T\n", new)
	fmt.Printf("Type of complex is %T\n", complex)
}
