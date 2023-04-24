package main

import "fmt"

func main() {
	s := "Hello, world!"
	p := &s

	fmt.Println(s, p)
	fmt.Println(*p)

	*p = "Hello, gopher!"

	fmt.Println(*p)
	fmt.Println(s)

	p = new(string)
	fmt.Println(p)
	fmt.Println(*p)
}
