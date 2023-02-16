package main

import "fmt"

// when a function executes it executes line by line
// and if you defer a line it will execute at the end of the function
//LIFO structure for the defered functions

func main() {
	fmt.Println("Hello world!")
	//this confirms the lIFO architecture
	defer fmt.Println("World1")
	defer fmt.Println("World2")
	fmt.Println("Surprie mother fuckers!!!")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
