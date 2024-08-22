package main

import "fmt"

func main() {
	n := 4
	edges := [][]int{
		{0, 1, 3},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}
	distanceThreshold := 4
	fmt.Printf("%d\n", findTheCity(n, edges, distanceThreshold))
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	return 0
}
