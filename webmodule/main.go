package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://iambishalchakraborty.github.io/"

func main() {
	fmt.Println("Welcome to WEB")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
