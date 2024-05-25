package main

import (
	"fmt"
	"sort"
)

func main() {

	// declarationa and initialisation
	var numbers = []int{2, 5, 3, 4, 1}
	fmt.Println(numbers)

	// appending value(s)
	numbers = append(numbers, 6, 7, 8);
	fmt.Println(numbers)

	// declaring with initialisation
	fruitList := make([]string, 5)
	fmt.Println(len(fruitList))

	// soring slice
	sort.Ints(numbers)
	fmt.Println(numbers)

	// removing an element from a slice based on index
	var index = 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println(numbers)

}