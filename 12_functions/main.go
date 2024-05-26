package main

import "fmt"

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func add(values ...int) int {
	sum := 0

	for _, val := range values {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(sum(5, 5))
	fmt.Println(add(1, 2, 3, 4, 5))
}