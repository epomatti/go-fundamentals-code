package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Interface
type printer interface {
	Print() string
}

// Object
type user struct {
	username string
	id       int
}

// Interface implementation
func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func main() {
	var p printer

	// User
	p = user{username: "adent", id: 42}
	fmt.Println(p.Print())

	// Menu Item
	p = menuItem{
		name: "Coffee",
		prices: map[string]float64{
			"small":  1.65,
			"medium": 1.80,
			"large":  1.95,
		},
	}
	fmt.Println(p.Print())

	// Type Assertion
	u, ok := p.(user)
	fmt.Println(u, ok)

	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	// Type Switch
	switch v := p.(type) {
	case user:
		fmt.Println("Found a user!", v)
	case menuItem:
		fmt.Println("Found a menu item!", v)
	default:
		fmt.Println("I'm not sure what this is...")
	}

}
