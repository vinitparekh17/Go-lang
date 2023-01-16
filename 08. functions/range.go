package main

import "fmt"

func main() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = (i + 1) + i
		fmt.Println(a[i])
	}

	// now intarating using range keyword
	for i, element := range a {
		fmt.Printf("index: %d, element: %d\n\n", i, element)
	}

	// let's find duplicate using range

	var b []int = []int{1, 2, 3, 4, 5, 6, 7, 3}

	for i, elm := range b {
		for j, elm2 := range b {
			if i > j && elm == elm2 {
				fmt.Println("Duplicate found: ", elm)
			}
		}
	}

	// easiest way to find duplicate

	for i, elm := range b {
		// in this approch we will never have i = j so little bit faster
		for j := i + 1; j < len(b); j++ {
			if elm == b[j] {
				fmt.Println("Duplicate found: ", elm)
			}
		}
	}
}
