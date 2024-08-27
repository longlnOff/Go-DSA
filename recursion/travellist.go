package main

import (
	"fmt"
)

// implement travel through list using recursion function
// in this example we calculate sum of all odd number
func sumOdds(index int, list []int) int {
	if index == len(list) {
		return 0
	}
	value := 0
	// check odd number
	if list[index] % 2 != 0 {
		value = list[index]
	}

	return sumOdds(index+1, list) + value
}

func productOfEvenNumbers(list []int) int {
	value := f(0, list)
	if value != 0 && value != 1 {
		return value
	} else {
		return 0
	}
}

func f(index int, list []int) int {
	if len(list) == 0 {
		return 0
	}
	if index == len(list) {
		return 1
	}
	value := 1
	if list[index] % 2 == 0 {
		value = list[index]
	}
	return f(index+1, list) * value
}

// https://www.youtube.com/watch?v=JBFFWnZIPx8&list=PLDV1Zeh2NRsCmu1lb9grUcljeYJtmgmYc&index=3
// another solution

func solution2(list []int) int {
	value := f1(0, list)
	if value == nil {
		return 0
	}
	return value.(int)
}
func f1(index int, list []int) any {
	if index == len(list) {
		return nil
	}

	product := f1(index+1, list)
	if list[index] % 2 != 0 {
		return product
	} else {
		if product == nil {
			return list[index]
		}
		return list[index] * product.(int)
	}

}

func main() {
	fmt.Println(productOfEvenNumbers([]int{1,3,5}))
	fmt.Println(solution2([]int{1,3,5}))

}