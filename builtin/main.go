package main

import (
	"fmt"

	// import package add
	"./add"
)

func main() {
	add.HelloWorld()
	fmt.Println(add.Add(10, 5))
	fmt.Println(add.Mul(10, 5))
}
