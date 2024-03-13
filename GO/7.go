package main

import "fmt"

func Seven() {
	// IS THIS DEEP CLONE??????
	matrix := [][]int{
		{1},
		{1},
		{1},
	}

	fmt.Println(matrix)

	matrix_clone := make([][]int, len(matrix))

	copy(matrix_clone, matrix)

	fmt.Println(matrix_clone)

	matrix_clone[0][0] = 2

	fmt.Println(matrix_clone, matrix)

	matrix_clone = append(matrix_clone, []int{1, 2, 3})
	fmt.Println(matrix_clone, matrix)
}