package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}

func main() {
	a := []int{1,3,5,5345,345,345,6,2,23}
	fmt.Println(insertionSort(a))
}

