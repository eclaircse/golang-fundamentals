package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the city:")

	// comma error

	input, _ := reader.ReadString('\n') // reads till it encounters '\n

	// _ ke jagah pe error hoga hain
	// since error is never used here toh _
	// it's kinda alternative to trycatch

	fmt.Println("The city name is:", input)
	fmt.Printf("Type of the city name is %T", input)

}