package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymetid=asjdnajsdnausdnuasd67867"

func main() {
	fmt.Println("Will be handling URLs now!")
	fmt.Println(myURL)
	//parsing the URL using URL library

	result, err := url.Parse(myURL) // we use the url class of the net package

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	queryParams := result.Query() // will store all the query params in a better way
	fmt.Println(queryParams)

	for _, val := range queryParams {
		fmt.Println(val)
	}

	// constructing a url string
	// always make sure you are passing a refrence and not the copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=Puttanpal",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("This is the another url", anotherUrl)

}
