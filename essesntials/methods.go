package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	rect := rectangle{10, 5}

	fmt.Println(rect)

	fmt.Println("area: ", rect.area())
	fmt.Println("perimeter: ", rect.perimeter())

	pointerToRect := &rect
	fmt.Println("area: ", pointerToRect.area())
	fmt.Println("perimeter: ", pointerToRect.perimeter())
}
