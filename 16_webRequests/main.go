package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.github.com/users/eclairjit"

func handleError(err error)  {
	if(err != nil) {
		panic(err)
	}
}

func main() {
	res, err := http.Get(url)

	handleError(err)

	defer res.Body.Close() // to close the connection when work's done

	// reading the response

	databytes, err := io.ReadAll(res.Body) // reading body (headers are also there)
	
	handleError(err)

	fmt.Println(string(databytes))

}