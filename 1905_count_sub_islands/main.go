package main

import "fmt"

func main() {
	// grid1 := [][]int{
	// 	{1, 1, 1, 0, 0},
	// 	{0, 1, 1, 1, 1},
	// 	{0, 0, 0, 0, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{1, 1, 0, 1, 1},
	// }

	// grid2 := [][]int{
	// 	{1, 1, 1, 0, 0},
	// 	{0, 0, 1, 1, 1},
	// 	{0, 1, 0, 0, 0},
	// 	{1, 0, 1, 1, 0},
	// 	{0, 1, 0, 1, 0},
	// }

	grid1 := [][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
	}

	grid2 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}

	fmt.Println(countSubIslands(grid1, grid2))
}

// Time complexity: O(m*n)
// Space complexity: O(m*n)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	count := 0
	m, n := len(grid2), len(grid2[0])

	arr := [][]int{}
	validIsland := false
	var (
		x, y int
		tile []int
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 {
				continue
			}

			arr = append(arr, []int{i, j})
			validIsland = true

			for len(arr) > 0 {
				tile = arr[0]
				arr = arr[1:]
				x, y = tile[0], tile[1]

				if grid2[x][y] == 1 {
					if x > 0 && grid2[x-1][y] == 1 {
						arr = append(arr, []int{x - 1, y})
					}

					if y < n-1 && grid2[x][y+1] == 1 {
						arr = append(arr, []int{x, y + 1})
					}

					if x < m-1 && grid2[x+1][y] == 1 {
						arr = append(arr, []int{x + 1, y})
					}

					if y > 0 && grid2[x][y-1] == 1 {
						arr = append(arr, []int{x, y - 1})
					}

					if grid1[x][y] != 1 {
						validIsland = false
					}
					grid2[x][y] = 0
				}
			}

			if validIsland {
				count += 1
			}
			validIsland = false
		}
	}

	return count
}
