package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string](float64)
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}

func (m *menu) addItem() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range data {
		if item.name == name {
			return errors.New("menu item already exists")
		}
	}

	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}

var in = bufio.NewReader(os.Stdin)

// Adds an item to the menu
func AddItem() error {
	return data.addItem()
}

// Print the menu for user interaction
func Print() {
	data.print()
}
