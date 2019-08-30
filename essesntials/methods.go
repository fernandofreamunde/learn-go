package main

import "fmt"

type rectangle struct {
	width, height int
}

func (this *rectangle) area() int {
	return this.width * this.height
}

func (this rectangle) perimeter() int {
	return 2*this.width + 2*this.height
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
