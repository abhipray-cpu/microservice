package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("This is how we handle get requests in golang")
	get1()
	postJSON()
	formData()
}

func get1() {
	const myUrl = "http://localhost:3000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("Unable to contact with machine")
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code of the req: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	//this is one of the way
	byteCount, _ := responseString.Write(content) // this is from a packer therefore will provide us with more options
	fmt.Println("This is the byteCOunt: ", byteCount)
	fmt.Println(responseString.String())
	// this is the second string
	if err != nil {
		fmt.Println("Unable to parse the response")
		panic(err)

	}
	// we are going to use string builders where we use get methods

	fmt.Println(string(content))

}
func post() {
	const url = "http://localhost:3000/post"
	// we send data generally in  json or url encoded form

}

// posting json data
func postJSON() {
	const url = "http://localhost:3000/post"
	fmt.Println("This is the JSON post function")
	//fake json payload
	//this is the new format in which we can create data in our case JSON data
	reqBody := strings.NewReader(`
	   {
          "courseName":"Let's go with golang",
		  "price":69,
		  "platform":"learncodeonline.in"
	   }
	`)
	res, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(string(content))
}

// poosting url encoded data
func formData() {
	const myurl = "http://localhost/3000/postform"
	fmt.Println("This is the URL encoded form function")
	//form data url
	data := url.Values{} // intially we create an empty key value pair
	// then we keep on adding values
	data.Add("firstname", "puttanpal")
	data.Add("lastname", "chaurasiya")
	data.Add("email", "dumkaabhipray@gmail.com")
	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(string(content))

}
