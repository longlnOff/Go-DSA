package sorting

import (
	"fmt"
	"testing"
)


func TestMergeSort(t *testing.T) {
	a := []int{1,3,4,5,675,3,34,34,3}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}


func MergeSort(a []int, left int, right int) {
	if left < right {
		m := left + (right-left) / 2
		MergeSort(a, left, m)
		MergeSort(a, m+1, right)
		Merge(a, left, m, right)
	}
}

func Merge(a []int, left int, m int, right int) {
	leftSize := m - left + 1
	rightSize := right - m
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)
	copy(leftArr, a[left:m+1])
	copy(rightArr, a[m+1:right+1])
	i := 0
	j := 0
	k := left
	for i < leftSize && j < rightSize {
		if leftArr[i] < rightArr[j] {
			a[k] = leftArr[i]
			i++
		} else {
			a[k] = rightArr[j]
			j++
		}
		k++
	}

	// copy the remaining element of two temporary array if any
	for i < leftSize {
		a[k] = leftArr[i]
		i++
		k++
	}
	for j < rightSize {
		a[k] = rightArr[j]
		j++
		k++
	}
}
