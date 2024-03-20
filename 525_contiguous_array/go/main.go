package main

import (
	"fmt"
)

// Submission result: 75ms (91.78%) | 6.75MB (99.93%)
func main() {
	nums := []int{0, 1, 0}
	fmt.Println(findMaxLength(nums))
}

func findMaxLength(nums []int) int {
	h := map[int]int{}
	h[0] = -1
	res := 0

	prefixSum := 0
	for i, n := range nums {
		if n == 0 {
			n = -1
		}
		prefixSum += n

		if x, ok := h[prefixSum]; !ok {
			h[prefixSum] = i
		} else {
			if i-x > res {
				res = i - x
			}
		}

	}

	return res
}
