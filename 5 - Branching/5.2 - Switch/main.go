package main

import "fmt"

func main() {

	switch i := 999; i {
	case 1:
		fmt.Println("firstcase")
	case 2 + 3, 2*i + 3:
		fmt.Println("second case")
	default:
		fmt.Println("default case")
	}

}
