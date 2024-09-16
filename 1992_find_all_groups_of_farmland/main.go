package main

import "fmt"

func main() {
	land := [][]int{
		{1, 0, 0}, {0, 1, 1}, {0, 1, 1},
	}
	res := findFarmland(land)
	for i := range res {
		fmt.Printf("%+v\n", res[i])
	}
}

func findFarmland(land [][]int) [][]int {
	res := [][]int{}
	m, n := len(land), len(land[0])

	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 { // the first time we found top left farmland tile, note the top left index
				bottom, right, top, left := i, j, i, j
				// look to the right until no farmland found
				for right < n-1 {
					if land[bottom][right+1] != 1 {
						break
					}
					right += 1
				}
				// look to the bottom until no farmland found
				for bottom < m-1 {
					if land[bottom+1][right] != 1 {
						break
					}
					bottom += 1
				}

				// now, we already find the bottom right tile, add to res
				res = append(res, []int{top, left, bottom, right})

				// clear all forest to farmland (change 1 to 0) from top left to bottom right
				for x := top; x <= bottom; x++ {
					for y := left; y <= right; y++ {
						land[x][y] = 0
					}
				}
			}
		}
	}

	return res
}
