package main

import "fmt"

var g int // global var init
func main() {
	var a int // local var init
	a = 10
	fmt.Println(a)
	g = 45
	fmt.Println(g)
}

// default value of a variable is 0
