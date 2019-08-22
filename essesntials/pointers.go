package main

import "fmt"

func zeroValue(iValue int) {
	iValue = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1

	fmt.Println("initial value of i: ", i)

	zeroValue(i)
	fmt.Println("Value of i after zeroValue(): ", i)

	zeroPointer(&i)
	fmt.Println("Value of i after zeroPointer(): ", i)

	fmt.Println("Value of the Pointer of i: ", &i)
}
