package main

import (
	"fmt"
	"math/rand" // just like python we can import explicit classes from a module
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("YOu got 1 bitch!")
	case 2:
		fmt.Println("YOu got 2 bitch!")
	case 3:
		fmt.Println("YOu got 3 bitch!")
		fallthrough // once this case is true other will alos be checked
	case 4:
		fmt.Println("YOu got 4 bitch!")
	case 5:
		fmt.Println("YOu got 5 bitch!")
	case 6:
		fmt.Println("YOu got 6 bitch!")
	case 7:
		fmt.Println("YOu got 7 bitch!")
	default:
		fmt.Println("bhal tohri mayi ka chodo")
	}
}
