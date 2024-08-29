package main

import "fmt"

func FindMatrixMax(matrix [][]int) (int, int) {
	return maxMatrixPrivate(0, 0, matrix)
}
func maxMatrixPrivate(row int, column int, matrix [][]int) (int, int) {
	if row == len(matrix) {
		return -1, -1
	}
	if column == len(matrix[0]) {
		return maxMatrixPrivate(row+1, 0, matrix)
	}
	row_max, col_max := maxMatrixPrivate(row, column+1, matrix)
	if row_max == -1 {
		row_max, col_max = row, column
	}
	if matrix[row][column] > matrix[row_max][col_max] {
		row_max, col_max = row, column
	}
	return row_max, col_max

}

func main() {
	a := [][]int{
		{0, 1, 2, 3},
		{4, 5, 99, 7},
		{4, 5, 999, 7},
		{4, 5, 99, 7},

	}
	fmt.Println(a)
	fmt.Println(FindMatrixMax(a))
}