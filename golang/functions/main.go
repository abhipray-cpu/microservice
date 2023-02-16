package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	puttanpal()
	mundal()
	fmt.Println("This is the output of the custom adder function", customAdder())
}

func puttanpal() {
	fmt.Println("Dhat tehri mayi ka chodo!")
	result := adder(69, 69)
	fmt.Println("This is the result of the function", result)
}

func mundal() {
	fmt.Println("Dhat tehri fuffi ka chodo!")
}

// need to define func signature as well
//
//	this is how we define a function in golang
//
// we need to define the type of function arguments and also the return type of the function if there is any
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// using the custom adder function in this we will take the user input
func customAdder() int64 {
	reader := bufio.NewReader(os.Stdin)
	num1, _ := reader.ReadString('\n')
	num2, _ := reader.ReadString('\n')
	num1_num, err1 := strconv.ParseInt(strings.TrimSpace(num1), 10, 64)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(num1_num)
	}
	num2_num, err2 := strconv.ParseInt(strings.TrimSpace(num2), 10, 64)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(num2_num)
	}
	return num1_num + num2_num

}
