package main

import "fmt"

func main() {
	// this is how we define a slice in golang
	days := []string{"Sunday", "Saturday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// classic way of looping through a low level language
	for day := 0; day < len(days); day++ {
		fmt.Println(days[day])
	}

	// this is how we loop through in modern high level languages
	for index, day := range days {
		fmt.Println(day, "is at ", index)
	}
	i := 1
	for i < 13 {
		if i == 3 {
			goto label1 // goto label_name labels are like functions that will execute the code inside them
		} else if i == 6 {
			goto label2
		} else if i == 9 {
			goto label3
		} else if i == 12 {
			goto label4
		}
		i++
	}
label1:
	fmt.Println("Teri maa ki chut hogayi!")
label2:
	fmt.Println("Teri fuffi ki chut hogayi!")
label3:
	fmt.Println("Teri apaa ki chut hogayi!")
label4:
	fmt.Println("Teri khala ki chut hogayi!")

}
