package main

import "fmt"

const Constant = "abcdefghijklmnopqrstuvwxyz" // first letter of variable if kept capital for global variables (not syntax but general practice)

func main() {
	var username string = "eclair"
	fmt.Println("Hello, " + username + "!")
	fmt.Printf("Variable is of type: %T \n", username)

	// Printf is to print formatted strings, doesn't print int
	// Print and Println are similar, only difference being, Println adds a newline but Print doesn't

	// go stores 0 in non-assigned integer and empty string in non-assigned string 

	var website = "https://www.google.com" // don't have to worry about data types, lexer takes care of that, but this variable can only store strings
	println(website)

	a := 1; // don't have to worry about type or var keyword, Go will infer it, but can only be used inside some method

	fmt.Println(a)
	fmt.Println(Constant)
}