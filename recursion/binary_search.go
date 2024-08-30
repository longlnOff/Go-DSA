package main

import "fmt"

func binarySeachRecursion(val int, arr []int) any {
	return binarySearchRecursivePrivate(0, len(arr)-1, val, arr)
}

func binarySearchRecursivePrivate(low int, high int, val int, arr []int) any {
	index := (low + high) / 2
	if low >= high+1 {
		return nil
	}
	var ret_val any
	if arr[index] < val {
		ret_val = binarySearchRecursivePrivate(index+1, high, val, arr)
	} else if arr[index] > val {
		ret_val = binarySearchRecursivePrivate(low, index-1, val, arr)
	} else {
		ret_val = index
	}

	return ret_val
}

func main() {
	a := []int{}
	fmt.Println(len(a))
	fmt.Println(binarySeachRecursion(7, a))
}