package main

import (
	"fmt"
	"math"
)

// declaration of the interface
type geometry interface {
	area() float64
	perimeter() float64
}

// declaration of the classes
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// to implement an interface in go we just need to implement all the methods in it
func (rect rectangle) area() float64 {
	return rect.width * rect.height
}
func (rect rectangle) perimeter() float64 {
	return 2*rect.height + 2*rect.width
}

func (circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}
func (circ circle) perimeter() float64 {
	return 2 * math.Pi * circ.radius
}

// if a variable have an interface type, then we can call all the methods in the named interface
func measure(geo geometry) {
	fmt.Println(geo)
	fmt.Println(geo.area())
	fmt.Println(geo.perimeter())
}

func main() {
	theRectangle := rectangle{3, 4}
	theCircle := circle{5}

	measure(theRectangle)
	measure(theCircle)
}
