package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {

	user := User{
		name:  "John",
		age:   25,
		email: "john@gmail.com",
	}

	fmt.Println(user)
	fmt.Printf("User: %+v\n", user) // for more detailed output

}