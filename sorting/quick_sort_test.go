package sorting

import (
	"fmt"
	"testing"
)


func TestQuickSort (t *testing.T) {
	a := []int{1,3,5,5345,345,345,6,2,23}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}


func QuickSort(a []int, low int, high int) {
	if low < high {
		pi := Partition(a, low, high)
		QuickSort(a, low, pi-1)
		QuickSort(a, pi+1, high)
	}

}

func Partition(a []int, low int, high int) int {
	pivot := a[high]
	i := low-1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i+1
}