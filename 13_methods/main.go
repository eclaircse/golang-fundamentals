package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

// methods
func (user User) greet() {
	fmt.Printf("Hello, I am %v.", user.name)
}

func main() {
	user := User{"John", 30, "john@gmail.com"}

	user.greet()
}