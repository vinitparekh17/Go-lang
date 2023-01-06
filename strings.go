package main

// import multiple packages
import (
	"fmt"
	"strings"
)

func main() {
	var a string = "Vinit"
	fmt.Println("Normal string: ", a)
	fmt.Printf("Hex bytes: ")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%x ", a[i])
	}
	fmt.Printf("\n%+q\n", a) // quoted string, %q is for single quotes and %+q is for double quotes

	// string functions
	fmt.Println("Length of string: ", len(a))

	fmt.Println("Contains: ", strings.Contains(a, "am"))
	fmt.Println("Count: ", strings.Count(a, "a"))
	msg := []string[]{"I", "am", "Vinit"} // array of strings
	fmt.Println("Concat:", strings.Join(msg, " "))
}
