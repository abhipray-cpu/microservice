package main

import "fmt"

func main() {
	languages := make(map[string]string) //thi is how we define a map make will not only create map but will also allocate memory
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS shorts for:", languages["JS"])

	// deleting values
	delete(languages, "RB")
	fmt.Println(languages)

	// looping through values
	for key, value := range languages {
		fmt.Println(key, value)
	}
}
