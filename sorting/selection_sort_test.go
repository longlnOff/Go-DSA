package sorting

import (
	"fmt"
	"testing"
)


func TestSelectionSort1 (t *testing.T) {
	a := []int{1,3,5,5345,345,345,6,2,23}
	fmt.Println(SelectionSort(a))
}


func SelectionSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i+1; j < len(a); j++ {
			if a[j] < a[i] {
				minIndex = j
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex]
	}

	return a
}