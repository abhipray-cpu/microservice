package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("This is JSON data generation")
	//EncodeJson()
	DecodeJson()
}

// encoding the JSON data
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "maakabhosda@1", []string{"web-dev", "reactJS"}},
		{" MERN Bootcamp", 399, "learncodeonline.in", "fuffikabhosda@1", []string{"web-dev", "reactJS", "express", "nodejs"}},
		{" MEVN Bootcamp", 499, "learncodeonline.in", "ammikabhosda@1", []string{"web-dev", "vueJS", "express", "nodejs"}},
		{" MEAN Bootcamp", 599, "learncodeonline.in", "khalakabhosda@1", []string{"web-dev", "angularJS", "express", "nodejs"}},
		{" Ruby Bootcamp", 699, "learncodeonline.in", "cuchikabhosda@1", []string{"web-dev", "ruby", "ruby on rails"}},
	}
	// packaging this data as JSON
	// we need to pass interface alternate of struct
	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// usig marshal indent
	// similar to prettier of mongo
	finalJson1, err1 := json.MarshalIndent(lcoCourses, "", "\t") // this is the format data,prefix,indent
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("%s\n", finalJson1)

}

// decoding the JSON data
func DecodeJson() {
	// json data that we receive will be in byte format
	// and back ticks are used for writing data in some custom format
	jsonData := []byte(`
	 {
		"coursename": " MEAN Bootcamp",
		"Print": 599,
		"website": "learncodeonline.in",
		"tags": [
				"web-dev",
				"angularJS",
				"express",
				"nodejs"
		]
}
   `)
	// verifying the correct JSON format

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("The json data was valid!")
		var lcoCourse course                 // so this var will be used to store the decoded data
		json.Unmarshal(jsonData, &lcoCourse) // we need to pass a refrence of the  interface where the unmarshalled data will be stored
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("The Json data was invalid!")
	}

	// we just want to add data to key value

	var myData map[string]interface{} // interface type since we are not sure what the type of the data we receive from net will be
	json.Unmarshal(jsonData, &myData) // we need to pass a refrence
	fmt.Printf("%#v\n", myData)

	for key, value := range myData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}
}

// in backticks we can create aliases
// - means that the field will not be reflected to who ever is consuming the api
type course struct {
	Name     string `json:"coursename"`
	Print    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this will not reflect to the user in output just like hiden in input type of html
	Tags     []string `json:"tags,omitempty"` // this will not be shown if the field is null that is no data
}
