// Description: Anonymous variable
//
// Anonymous variable is a variable that is declared but not used in the program
// Anonymous variable is declared using the _ operator in golang
// Anonymous variable is used to ignore the value of a variable

// How to use anonymous variable?
// 1. when you want to ignore the value of a variable
// 2. when you want to ignore the index of a variable
// 3. when you want to ignore the error of a function
// 4. when you want to ignore the return value of a function

package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 1, 3}

	for _, element := range a {
		fmt.Printf("%d\n", element)
	}
}
