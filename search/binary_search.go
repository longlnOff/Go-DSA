package main

import "fmt"

// iterative approach
func binarySearchIterative(val int, arr []int) any {
	low := 0
	high := len(arr)
	for low <= high {
		index := (low + high) / 2
		if arr[index] < val {
			low = index + 1
		} else if arr[index] > val {
			high = index - 1
		} else {
			return index
		}
	}

	return nil
}

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
	a := []int{1,3,4,5,6,7,8,9,10}
	fmt.Println(len(a))
	fmt.Println(binarySearchIterative(4, a))
	fmt.Println(binarySeachRecursion(4, a))
}