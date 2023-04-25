package main

import "fmt"

func main() {
	result, ok := divide(1, 2)
	if ok {
		fmt.Println(result)
	}
}

func divide(l, r int) (result int, ok bool) {
	if r == 0 {
		return // 0, false
	}
	result = l / r
	ok = true
	return
}
