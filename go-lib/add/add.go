package main

import "C"

func main() {}

//export add
func add(a, b C.int) C.int {
	return a + b
}
