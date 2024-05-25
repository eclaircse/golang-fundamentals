package main

import "fmt"

func main() {
	var numbers [5]int // array declaration
	numbers[0] = 1

	fmt.Println(numbers)

	var fruits = [3]string {"apple", "orange", "banana"}

	fmt.Println(fruits)

	fmt.Println(len(fruits)) // prints length of the array
}