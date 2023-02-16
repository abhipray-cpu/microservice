package main

import "fmt"

func main() {
	// there is no inheritance in golang
	abhipray := User{"Abhipray", "dumkaabhipray@gmail.com", true, 23}

	fmt.Print(abhipray)

	fmt.Printf("Abhipray  detaials are: %+v\n", abhipray)
	fmt.Printf("Name is: %v and email is %v", abhipray.Name, abhipray.Email)
}

// name should be in capital since the fields will be public
// structs are similar to constructors in C++
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
