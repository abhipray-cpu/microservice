package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files in golang")
	content := "This needs to go in file by puttanpal copies code"

	file, err := os.Create("./this_is_it.txt") // we create file using os package

	checkNilErr(err)

	length, err := io.WriteString(file, content) // we are using io to write to the file
	checkNilErr(err)
	fmt.Println("length is :", length)

	defer file.Close() // we do this to ensure that the file is closed at last when we are done with the operations.
	reader("./this_is_it.txt")

}

func reader(fileName string) {
	// os is used for creating the files
	//ioutil is used for manipulating the file

	dataBytes, err := ioutil.ReadFile(fileName) // we are using ioutil to read the file
	checkNilErr(err)
	// the file is not read in string format but in bytes format

	fmt.Println("Text data inside the file is \n", dataBytes)
	fmt.Println("Text data inside the file is \n", string(dataBytes))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err) //panic is used to raise an error just like alert
	}
}
