package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(123)
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
