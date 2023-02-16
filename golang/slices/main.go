// slices is bascially arrays on steroids

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("let's play with slices!!")
	var fruitList = []string{"Apple", "Tomato", "Peach"} //this is how we define slice
	fmt.Printf("Type of fruitlist is %T\n", fruitList)   // %T operator is used to get the type of variable
	// this is how we insert data to slices
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // just normal slicing of array
	fmt.Println((fruitList))

	highScores := make([]int, 4) // 2nd way of defining slices here we are using make and is reserving 4 blocks of mem

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 567
	//highScores[4] = 7777
	// we can append but we can't assign using index that is out of bound
	highScores = append(highScores, 55555, 7666, 77777)
	fmt.Println(highScores)

	fmt.Print(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// removing a value from slice based on index

	var courses = []string{"mkb", "bc", "mc", "bkl", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
