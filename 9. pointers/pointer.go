package main

import "fmt"

// pointer is a variable that stores the address of another variable
// pointer var is declared using the * operator
// syntax of pointer:
// var var_name *type
// var_name = &var // & is the address of operator and it returns the address of the variable just like c and c++
func main() {
	var a int = 10
	var b *int = &a

	fmt.Println(b)
	fmt.Printf("%T\n", b)  // type of b is *int which means it is a pointer to an int
	fmt.Printf("%d\n", *b) // displays the value at the address stored in b

	// pointer arithmetic is not allowed in go
	// b++ // this will give an error

	c := 54
	var ptr *int = &c

	fmt.Println(ptr)  // prints memory address of c
	fmt.Println(*ptr) // prints 54

	// pointer to a pointer
	var ptrptr **int = &ptr
	fmt.Println(ptrptr) // prints memory address of ptr

	*ptr = 100
	fmt.Println(c) // prints 100

	// pointer to a struct

	type person struct {
		name    string
		age     int
		hobboes []string
	}

	p := person{
		name:    "Vinit",
		age:     20,
		hobboes: []string{"coding", "reading", "gaming"},
	}

	var ptrp *person = &p

	fmt.Println(*ptrp) // prints memory address of p
}
