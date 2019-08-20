package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice with setted values: ", s)
	fmt.Println("slice value: ", s[0])
	fmt.Println("slice lenght: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice with new values: ", s)
	fmt.Println("new slice lenght: ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("clonned/copied slice: ", c)

	l := s[2:5]
	fmt.Println("slice of slice: ", l)

	l = s[:5]
	fmt.Println("slice up to the 5th element(excluding)", l)

	l = s[2:]
	fmt.Println("sliced slice up from 2thd element(including)", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declared and initialized inline slice:", t)

	md := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		md[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			md[i][j] = i + j
		}
	}
	// nice read: https://blog.golang.org/go-slices-usage-and-internals
	fmt.Println("multi dymentional slice: ", md)
}
