package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println(divide(10, 0))

	// result, err := divide1(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("result: ", result)

	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result: ", result)

}

//lint:ignore U1000 just for testing
func divide(l, r int) int {
	return l / r
}

//lint:ignore U1000 just for testing
func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")

	}
	return l / r, nil
}

// Using named variables
func divide2(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
