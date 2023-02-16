package main

import "fmt"

func main() {
	loginCount := 23
	var result string
	// see the indetation and the spacing this needs to be defined in the same exact order or else an error will be thrown
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >= 10 && loginCount < 20 {
		result = "Bhadwa salaa!"
	} else {
		result = "Randwa sala"
	}
	fmt.Println("This is the message for you bitvh!! \n", result)
}
