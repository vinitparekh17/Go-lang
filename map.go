// Description: Map in Go
//
// Map is a data structure that stores key-value pairs in a random order or user defined order
// Map is a reference type in Go
// Map is a hash table in Go
// Map is a collection of unordered key-value pairs which is also know in other languages as dictionary or associative array

package main

import "fmt"

func main() {
	a := map[string]int{
		"Vinit": 1,
		"Raj":   2,
		"Chiku": 3,
	}

	for key, value := range a {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}
