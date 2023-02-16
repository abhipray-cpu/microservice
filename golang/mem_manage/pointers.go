package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers!")
	var ptr1 *int // * is of data type pointer
	fmt.Println("The value of pointer is", ptr1)
	myNumber := 69
	fmt.Println("The address is", &myNumber)
	var ptr = &myNumber
	fmt.Println("The pointer is referencing to", ptr)
	fmt.Println("The value to which pointer is referencing is", *ptr)
	*ptr = *ptr + 2
	fmt.Println(myNumber)

}
