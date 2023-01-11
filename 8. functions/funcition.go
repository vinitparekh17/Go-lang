package main

import "fmt"

func main() {

	fmt.Println(myfun(1, 2))
	fmt.Printf("%T", stringInt)
}

func myfun(a, b int) int {
	return a + b
}

// diff type of perameters taken
func stringInt(a int, b string) (string, int) {
	return b, a
}

// function in go
// user defined function is declred outside the main function
// function name should start with a capital letter if it is to be exported
// function can return multiple values

// syntax:
// func function_name(parameter_name type) return_type {
// 	// function body
// }
