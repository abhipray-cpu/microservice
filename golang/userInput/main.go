package main

import (
	"bufio" // this mod is used for declaring a buffer which will be used for taking input
	"fmt"
	"os" // our general OS module
)

// buffio package is a buffer that will allow to read from user input
func main() {
	welcome := "This is a simple welcom message lawdiya!!"
	fmt.Println(welcome)

	// this is how we take user input by creating a buffer
	reader := bufio.NewReader(os.Stdin) // this is how we create a read buffer
	fmt.Println("NIE ki maa ki chut kitni baar?")

	// comma ok syntax || error ok syntax
	// we don't have try catch in go
	// walrus since the user input can be of unpredictiable type
	input, _err := reader.ReadString('\n') // \n is an ender
	// we can use this input and _err as try and catch blocks
	fmt.Println(input)
	fmt.Println(_err)
	// the type of user input will be a string

}
