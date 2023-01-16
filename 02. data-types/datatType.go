// Data Type in Golang is as follows:
// bool types
// includes : true / false

// string types

// derived types
// (a) Pointer types,
// (b) Array types,
// (c) Structure types,
// (d) Union types,
// (e) Function types,
// (f) Slice types,
// (g) Interface types,
// (h) Map types,
// (i) Channel Types

// All numeric types : a) int b) float

// int  int8  int16  int32  int64
// int used to store integer of any size
// int8 used to store 8 bit integer
// int16 used to store 16 bit integer
// int32 used to store 32 bit integer
// int64 used to store 64 bit integer

// What is unsigned integer?
// unsigned integer is a integer which can't be negative

// uint uint8 uint16 uint32 uint64 uintptr
/// uintptr is used to store the uninterpreted bits of a pointer value

// what is byte?
// byte is alias of uint8 which is used to store 8 bit unsigned integer

// what is rune?
// rune is alias of int32 which is used to store 32 bit integer

// float32 float64
// float32 is used to store 32 bit floating point number
// float64 is used to store 64 bit floating point number

// What is complex number in golang?
// complex number is a number which has real and imaginary part
// for example 2 + 3i

// complex64 complex128
// complex64 is used to store 32 bit complex number
// complex128 is used to store 64 bit complex number
package main

import "fmt"

func main() {
	var d int
	d = 5
	y := 6
	x := "Vinit"
	// := is used to dynamic type decleration
	fmt.Println(d)
	fmt.Printf("Type of y is %T\n", x)
	fmt.Println(y)
	fmt.Printf("Type of y is %T", y)
	// printf used to get output in format , here %T displays int whereas PrintLn takes
}
