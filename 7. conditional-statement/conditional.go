// tota; 5 conditional statements

// 1: if

// 2: if.. else

// 3: Nested if

// 4: switch

// 5: select
// select statement is used to choose from multiple send/receive channel operations.

// diff between switch and select:

// switch is used to select from many alternatives in if-else statements.
// select is used to select from many send/receive channel operations.

// 6: for

// keyword: break, continue, goto

// how to use goto and label in go
// write a program that uses goto and label.
// for example: if i == 7 goto label

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}

	j := 0

label:
	fmt.Println(j)
	j++
	if j < 10 {
		goto label // this will cause a loop
	}
}
