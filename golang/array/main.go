package main

import "fmt"

func main() {
	fmt.Println("bhak tehri maayii ka chodo!")

	var fruitList [4]string // the difference betweeen slice and array is that in slice we don't define the size

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Strawberry"
	fruitList[3] = "Melon"
	fmt.Println("Fruit list is", fruitList)
	fmt.Println("The number of fruits we have is ", len(fruitList))

	var veggies = [3]string{"mushroom", "sweet corn", "peas"}
	fmt.Println("Veggies list is ", veggies)
	fmt.Println("We have ", len(veggies), " number of vegetables in my basket")
}
