package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

// we will be using netHTTP package for handling web requests
func main() {
	fmt.Println("will be making a web request")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The response of the type %T \n", res)
	fmt.Println(res)

	defer res.Body.Close() // caller's responsibilty to close the connection
	// majority of the reading is done using ioUtil

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the response", string(result))

}
