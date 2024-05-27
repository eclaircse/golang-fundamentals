package main

import "fmt"

func main() {
	defer fmt.Print("1234")
	defer fmt.Println(" world!")
	fmt.Print("Hello")
}

// prints
// Hello world!
// 1234
// defer statements are executed just before the funcion
// returns, and that too in the reverse order in
// which they were declared (LIFO)