package main

import "fmt"

func main() {

	// initialize a map
	languages := make(map[string]string)

	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"

	fmt.Println(languages)
	fmt.Println(languages["en"])

	// delete a key
	delete(languages, "fr")
	fmt.Println(languages)

	// print all values
	for _, value := range languages {
		fmt.Println(value)
	}

}