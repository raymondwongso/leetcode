package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
		{3, 3},
	}
	fmt.Println(xorQueries(arr, queries))
}

// time complexity: O(max(n, m)) where n is len of arr, m is len of queries
// space complexity: O(m) for res array
func xorQueries(arr []int, queries [][]int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	res := make([]int, len(queries))
	for i := range res {
		left, right := queries[i][0], queries[i][1]
		if left == 0 {
			res[i] = arr[right]
		} else {
			res[i] = arr[right] ^ arr[left-1]
		}
	}

	return res
}
