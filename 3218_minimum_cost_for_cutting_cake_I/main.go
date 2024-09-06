package main

import (
	"fmt"
	"sort"
)

func main() {
	m, n := 3, 2
	horizontalCut := []int{1, 3}
	verticalCut := []int{5}
	fmt.Println(minimumCost(m, n, horizontalCut, verticalCut))
}

// time complexity: O(n)
// space complexity: O(1)
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)
	hindex, vindex := len(horizontalCut)-1, len(verticalCut)-1
	hmulti, vmulti := 1, 1

	cost := 0
	for hindex >= 0 || vindex >= 0 {
		if vindex < 0 {
			cost += horizontalCut[hindex] * hmulti
			hindex -= 1
			continue
		}

		if hindex < 0 {
			cost += verticalCut[vindex] * vmulti
			vindex -= 1
			continue
		}

		if horizontalCut[hindex] >= verticalCut[vindex] {
			cost += horizontalCut[hindex] * hmulti
			vmulti += 1
			hindex -= 1
		} else {
			cost += verticalCut[vindex] * vmulti
			hmulti += 1
			vindex -= 1
		}
	}

	return cost
}
