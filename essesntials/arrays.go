package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[3] = 29
	fmt.Println("with one value set: ", a)
	fmt.Println("the value set: ", a[3])

	fmt.Println("array length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared array: ", b)

	var twoDimentional [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimentional[i][j] = i + j
		}
	}
	fmt.Println("2dimentional array: ", twoDimentional)
	fmt.Println("2dimentional value: ", twoDimentional[1][2])
}
