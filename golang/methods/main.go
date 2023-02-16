package main

import "fmt"

func main() {
	// there is no inheritance in golang
	abhipray := User{"Abhipray", "dumkaabhipray@gmail.com", true, 23}

	fmt.Print(abhipray)

	fmt.Printf("Abhipray  detaials are: %+v\n", abhipray) // %+v operator is used to print in a better way
	fmt.Printf("Name is: %v and email is %v\n", abhipray.Name, abhipray.Email)
	abhipray.GetStatus() // now we can directly use these methods with the struct object we created
	abhipray.NewMail()
}

// name should be in capital since the fields will be public
// in golang public features are defines usign Caps for first letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

	// first letter capital are exportable
	// and small caps are private
}

// here the functions are accepting structs as data type therefore we can use them with structs values
// this is how we define methods the syntax is a little different from functions
func (u User) GetStatus() {
	fmt.Println(u.Name, "is active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "puttanpal@horny.com"
	fmt.Println("Email of", u.Name, "is", u.Email)
}
