package main

import "fmt"

func main() {
	//make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("value 1: ", v1)
	fmt.Println("map length: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	// clears up if the key exists or not
	// in this case it returns false
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	//inline init
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("another map: ", n)
}
