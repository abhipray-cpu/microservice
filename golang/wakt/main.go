package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to study of time in golang")
	presentTime := time.Now() // walrus operator is used for type independent allocation
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating time from input values
	createDate := time.Date(2020, time.September, 10, 19, 45, 7, 0, time.UTC)
	fmt.Println("Today's date is ,", createDate)

}
