package main

// Goto statement rules
// 1. Can leave block
// 2. Can jump to containing block
// 3. Can't jump aftger variable declaration
// 4. Can't jump into another block

func main() {
	myFunc()
}

func myFunc() {
	i := 10
	if i < 15 {
		goto myLabel
	}
myLabel:
	j := 42
	for ; i < 15; i++ {
		println(j)
	}
}
