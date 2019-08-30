package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("No 42 Allowed!")
	}

	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 7 {
		return -1, &argError{ arg, "Can't work with this"}
	}

	return arg + 3, nil
}

func main() {
	fmt.Println("hello there!")

	r, e := f1(42)
	fmt.Println("f1 with 42 - ", r, " error - ", e)
	r, e = f1(43)
	fmt.Println("f1 with 43 - ", r, " error - ", e)

	r, e = f2(7)
	fmt.Println("f2 with 7 - ", r, " error - ", e)
	r, e = f2(8)
	fmt.Println("f2 with 8 - ", r, " error - ", e)

	fmt.Println("\n\n\nand now as in the example in the site:")
	
	for _, i := range []int{7,42} {

		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 works: ", r)
		}
	}

	//If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion
	fmt.Println("\n\n----------------")
	_, d := f2(7)
	if ae, ok := d.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	
	fmt.Println("\n\n----------------")
	_, f := f2(42)
	if ae, ok := f.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
