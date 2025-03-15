  package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(10)
	fmt.Println("Your random number is: ", randomNumber)
}