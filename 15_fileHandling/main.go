package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleError(err error)  {
	if(err != nil) {
		panic(err)
	}
}

func readFile(fileName string)  {
	databyte, err := ioutil.ReadFile(fileName)
	handleError(err)

	fmt.Println("Text inside the file: ", string(databyte))
}

func main() {

	file, err := os.Create("./test.txt")

	handleError(err)

	defer file.Close()

	length, err := io.WriteString(file ,"Hello, World!")

	handleError(err)

	fmt.Println("Length: ", length)

	readFile("./test.txt")

}