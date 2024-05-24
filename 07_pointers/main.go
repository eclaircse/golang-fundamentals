package main

import "fmt"

func main() {
	num := 11

	var ptr = &num

	fmt.Println("Value of num is: ", *ptr) // prints 11
	fmt.Println("Address of num is: ", ptr) // prints the address

	*ptr = *ptr + 10

	fmt.Println("Value of num is: ", num) // prints 21
}