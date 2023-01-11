package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	createFile()
	readFile("./data.txt")
}

func createFile() {
	content := "This line is supposed to be in the file -Code by Vinit Parekh"
	file, err := os.Create("./data.txt")

	checkNilErr(err)

	if file != nil {
		// creating a file is the os operation so using os package
		len, err := io.WriteString(file, content)

		checkNilErr(err)

		if len != 0 {
			fmt.Printf("File has been created with the given content!\n")
			fmt.Println("Length of file is: ", len)
			defer file.Close()
		}
	}
}

func readFile(filePath string) {
	// io utils used to read file
	data, err := ioutil.ReadFile(filePath)

	checkNilErr(err)

	// fmt.Println("Content of the file is as follow as:\n", data)
	// to see this data bytes in readable form pass it in string func
	fmt.Println(string(data))
}

// err nil handler to avoid repeatative stuff - (code)
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
