package main

import (
	"fmt"
	"os"
	"strconv"
)

// So one BS inside of another with some logic difference in finding the right row
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rowStart, rowEnd := 0, len(matrix)-1

	// Loops through the rows
	for rowStart <= rowEnd {
		rowMid := (rowStart + rowEnd) / 2

		if target < matrix[rowMid][0] {
			rowEnd = rowMid - 1
		} else if target > matrix[rowMid][len(matrix[rowMid])-1] {
			rowStart = rowMid + 1
		} else {
			colStart, colEnd := 0, len(matrix[0])-1
			for colStart <= colEnd { // loop through the rows cols using BS
				colMid := (colStart + colEnd) / 2
				if matrix[rowMid][colMid] == target {
					return true
				} else if matrix[rowMid][colMid] < target {
					colStart = colMid + 1
				} else {
					colEnd = colMid - 1
				}
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
		{66, 67, 73, 78},
		{80, 86, 93, 99},
	}
	target, _ := strconv.Atoi(os.Args[1])
	fmt.Println(searchMatrix(matrix, target))
}
