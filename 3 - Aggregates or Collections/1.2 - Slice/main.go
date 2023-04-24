package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(s)
	fmt.Println(s[1])

	s[1] = "Chai Tea"
	fmt.Println(s[1])

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)

	slices.Delete(s, 1, 2)

	fmt.Println(s)

	// arr2 := s

	// arr2[2] = "Chai Latte"
	// fmt.Println(s, arr2)
}
