package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNumber)
}