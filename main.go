package main

import "fmt"

func main() {
	i := 42

	p := &i         // point to i
	fmt.Println(*p) // get value from p
	*p = 21
	fmt.Println(i) // new value of i
}
