package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	m, n := 3, 1
	fmt.Println(construct2DArray(original, m, n))
}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)
	for i := range m {
		start := i * n
		end := start + n
		res[i] = original[start:end]
	}

	return res
}
