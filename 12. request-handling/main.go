package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://lco.dev")
	fetal(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fetal(err)
	fmt.Println(string(data))
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
