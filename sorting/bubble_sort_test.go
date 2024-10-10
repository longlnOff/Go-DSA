package sorting

import (
	"fmt"
	"testing"
)


func TestBubbleSort(t *testing.T) {
	a := []int{1,3,4,4,4,6,23,34,5,3,3,4123,4234,34}
	fmt.Println(BubbleSort(a))
}

func BubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		swapped := false
		for j := 0; j < len(a)-1; j++ {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
				swapped = true
			}
		}
		if !swapped {
			break	
		}
	}

	return a
}