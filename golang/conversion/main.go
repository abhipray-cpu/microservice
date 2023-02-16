package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Plz rate our pizza between 1 and 5")
	/*
	   The input which we read via a buffer is of type string
	*/
	reader := bufio.NewReader(os.Stdin) // we read input by initializing a buffer
	input, _ := reader.ReadString('\n') // \n here is the delimetre at which we stop reading the input
	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // the type of buffer reader is always string
	// therefore we need to convert the string value to our required data type

	// this means there is some error and is basically try and catch alternative of golang
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}
	// bikris chod else should start just after } of if like in ejs

}
