package main

import "fmt"

func main() {
	// Function
	var i int
	a := isEvenFunction(i)
	fmt.Println(a)

	// Method
	var mi myInt
	b := mi.isEvenMethod()
	fmt.Println(b)
}

// Function
func isEvenFunction(i int) bool {
	return i%2 == 0
}

// Method
type myInt int

func (i myInt) isEvenMethod() bool {
	return int(i)%2 == 0
}
