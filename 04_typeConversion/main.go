package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number:")
	input, _ := reader.ReadString('\n')
	num, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(error!=nil) {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Value of number plus 1: ", num+1)
	}
}