package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	_, err := toUpper("")
	if err != nil {
		fmt.Println(fmt.Errorf("invalid output: %w", err))
	}
}

func toUpper(a string) (string, error) {
	if len(a) == 0 {
		return "", errors.New("a string cannot be empty")
	}
	return strings.ToUpper(a), nil
}
