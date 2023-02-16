package main

import "fmt"

/*
these walrus operators are not allowed outside functions
jwtToke := "nama tulla nama tulla"
*/

// definine the constants
const LoginToken string = "puttanpal will fuck the queen in hell"

// first letter cap is equivalent to public resource
func main() {
	var user string = "puttanpal" // this variable needs to be used
	// or else an error will be thrown
	fmt.Println(user)
	// %T this is a placeholder for getting the tpe of the variable
	fmt.Printf("Variable is of type: %T \n", user) //fp is a shortcut

	var isLoggedIN bool = false
	fmt.Println(isLoggedIN)
	fmt.Printf("Variable is of type: %T \n", isLoggedIN) //fp is a shortcut

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal) //fp is a shortcut

	var smallFloat float32 = 255.45533423232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat) //fp is a shortcut

	var smallFloat1 float64 = 255.45533423232
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type: %T \n", smallFloat1) //fp is a shortcut

	// default values and some aliases

	var anotherVariable int64
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type  %T \n", anotherVariable) //fp is a shortcut

	// implcit type
	var website = "dhat teri maa ki chut" // lexer will automatically decide the type and it will be fixed
	fmt.Println(website)

	num_of_users := 69699 // this is a flexible way can use it when we are not sure about the type
	// of input we will get
	fmt.Println(num_of_users)

	fmt.Println(LoginToken)
	fmt.Printf("This is the type : %T \n", LoginToken)
}
