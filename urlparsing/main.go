package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.amazon.in/Apple-iPhone-13-512GB-Blue/dp/B09G9JJT7M/ref=sr_1_7?crid=1SXXXGGVB6O33&keywords=iphone%2B13&qid=1651115777&sprefix=iphone%2B1%2Caps%2C650&sr=8-7&th=1"

func main() {
	fmt.Println("URL")

	result, _ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	//fmt.Println(qparams)
	//fmt.Println(qparams["keywords"])

	for s, strings := range qparams {
		fmt.Printf("%v : %v\n", s, strings)
	}

	createurl := &url.URL{
		Scheme: "http",
		Host:   "www.amazon.in",
		Path:   "Apple-iPhone-13-512GB-Blue/dp/B09G9JJT7M/ref=sr_1_7",
	}

	anotherurl := createurl.String()
	fmt.Println(anotherurl)

}
