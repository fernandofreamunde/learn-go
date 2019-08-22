package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 25})

	fmt.Println(person{"Fred", 0})

	fmt.Println(&person{"Ann", 40})

	s := person{"Sean", 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age++
	fmt.Println(sp.age)
}
