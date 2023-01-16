package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://www.google.com:43/search?q=go+programming+language&page=2"

func main() {

	result, err := url.Parse(URL)
	fetal(err)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queries := result.Query()

	for _, value := range queries {
		fmt.Println(value)
	}
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
