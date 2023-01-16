// syntax of array
// var array_name [size]type

// mathods of array
// len(array_name) // returns the size of the array
// array_name[index] // returns the value at the index
// array_name[index] = value // sets the value at the index
// array_name[:] // returns the whole array
// array_name[start_index:end_index] // returns the array from start_index to end_index-1

// types of array: 2d array, 3d array, jagged array
// 2d array: var array_name [size1][size2]type
// 3d array: var array_name [size1][size2][size3]type
// jagged array: var array_name [size1][]type

// what is jagged array?
// jagged array is an array of arrays of different sizes and types, in short nested arrays

package main

import (
	"fmt"
)

func main() {
	// 1d array
	var a [5]int

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println("1d Arr", a)

	// 2d array
	var b [2][3]int

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = i + j
		}
	}

	fmt.Println("2d Arr", b)

	// 3d array
	var c [2][3][4]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			for k := 0; k < len(c[i][j]); k++ {
				c[i][j][k] = i + j + k
			}
		}
	}

	fmt.Println("3d Arr", c)

	// jagged array
	var d [4][]int
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, i+1)
		for j := 0; j < len(d[i]); j++ {
			d[i][j] = i + j
		}
	}

	// explaination of the above code
	// jagged array is formed like a matrix
	// it contains 4 rows and 4 columns
	// the first row contains 1 element [0]
	// the second row contains 2 elements [1, 2]
	// the third row contains 3 elements [2, 3, 4]
	// the fourth row contains 4 elements [3, 4, 5, 6]

	// columns are not fixed in jagged array
	// rows are fixed

	// how to definded columns in jagged array?
	// d[i] = make([]int, i+1)
	// i+1 is the number of columns in the ith row
	// make is a built in function that is used to create slices, maps, and channels only

	fmt.Println("Jagged Arr", d)
	// declare a dynamic type jagged array without loop
	var e [3][]interface{}
	// interface{} is a type that represents any type
	// it is used to create a dynamic type array
	e[0] = []interface{}{1, "a", 3.14}
	e[1] = []interface{}{1, "a", "b"}
	e[2] = []interface{}{1, "a", 3.14, "b", 5, "c"}

	fmt.Println("Jagged Arr", e)
}
