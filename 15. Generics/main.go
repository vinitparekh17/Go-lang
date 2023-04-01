package main

import "fmt"

// A function that takes array of key value pair where key is string and value is integer
func sumOfnum(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// A function that takes array of key value pair where key is string and value is float
func sumofFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func Genericzfun[Key comparable, Value int64 | float64](m map[Key]Value) Value {
	var s Value
	for _, ans := range m {
		s += ans
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumOfnum(ints), sumofFloat(floats))
	fmt.Printf("Generic Sums: %v and %v\n", Genericzfun(ints), Genericzfun(floats))
}
