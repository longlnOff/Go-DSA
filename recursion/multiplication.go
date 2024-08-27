package main

import "fmt"

// implement multiplication using recursion

func multiplication(a int, b int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return multiplication(a+1, b) - b
	} else {
		return multiplication(a-1, b) + b
	}
}

func main() {
	fmt.Println(multiplication(-4, 5))
}