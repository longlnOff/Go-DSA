package main

import "fmt"

// find max value in list and return that value and its index
func findMax(list []int) (any, any) {
	return findMaxPrivate(0, list)
}
func findMaxPrivate(index int, list []int) (any, any) {
	if len(list) == index {
		return nil, nil
	}
	i, v := findMaxPrivate(index+1, list)
	if i == nil || v.(int) < list[index] {
		return index, list[index]
	} else {
		return i.(int), v.(int)
	}
}

func main() {
	a, b := findMax([]int{})
	fmt.Println(a, b)
}