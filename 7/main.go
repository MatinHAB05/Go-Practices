package main

import (
	"fmt"
	"sync"
)

func multiplyRow(a [][]int, b [][]int, result [][]int, row int, wg *sync.WaitGroup) {
	defer wg.Done()

	colsB := len(b[0])
	colsA := len(a[0])

	for j := 0; j < colsB; j++ {
		sum := 0
		for k := 0; k < colsA; k++ {
			sum += a[row][k] * b[k][j]
		}
		result[row][j] = sum
	}
}

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	B := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	rowsA := len(A)
	colsB := len(B[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	var wg sync.WaitGroup

	for i := 0; i < rowsA; i++ {
		wg.Add(1)
		go multiplyRow(A, B, result, i, &wg)
	}

	wg.Wait()

	fmt.Println("Result:")
	for _, row := range result {
		fmt.Println(row)
	}
}
