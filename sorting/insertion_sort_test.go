package sorting

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{1,5,6,4,5,7,8,3,23,2}
	fmt.Println(InsertionSort(a))
}

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}