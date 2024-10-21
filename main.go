package main

import "fmt"

type Pos struct {
	x, y int
}

func main() {
	fmt.Println(Pos{5, 3})

	point := Pos{1, 2}
	fmt.Println(point.x, point.y)
	point.x = 10
	fmt.Println(point.x, point.y)

	p := &point
	p.x = 5
	fmt.Println(point.x, point.y)

	axisPoint := Pos{x: 1} // y is unset => y = 0
	fmt.Println(axisPoint.x, axisPoint.y)
}
