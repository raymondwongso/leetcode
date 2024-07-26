package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/sort-the-jumbled-numbers/
func main() {
	// mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums := []int{991, 338, 38}

	fmt.Printf("%+v\n", sortJumbled(mapping, nums))
}

/*
Intuition:
Transform num from nums according to mapping
Create hashmap, using transformed num as key and array of original num as value
For each transformed num, append to the array value of the current original num
Sort the hashmap by its key in asc manner. For each key, append to the result

Time complexity: O(n log n)
  - Iterating nums, where n is len of nums, takes O(n).
    Inside nums iteration, iterate again for each digits (d) of num
    because num is at most 10^9 (int32), d is at most 10.
    compared 10 with int32.Max, we can conclude d as O(1) compared to n.
  - Sorting Keys
    sort.Ints use pattern-defeating quicksort (pdq), O(n log n)
  - Iterating key in keys, taes O(n)

So total is O(n + n log n + n) = O(n log n)
Space Complexity: O(n)
- worst case for m: distinct trasnformed num for each nums. O(n)
- worst case for res: store the original nums, O(n)
- worst case for keys: distinct transformed num for each nums. O(n)
So total is O(n)
*/
func sortJumbled(mapping []int, nums []int) []int {
	m := map[int][]int{}
	div, mod, mult := 1, 10, 1
	res := []int{}
	keys := []int{}

	for _, num := range nums {
		mappedNum := 0
		// calculating the transformed num
		for {
			mappedNum += mapping[num/div%mod] * mult
			div *= 10
			mult *= 10
			if div > num {
				break
			}
		}
		div = 1
		mult = 1

		if _, ok := m[mappedNum]; !ok {
			keys = append(keys, mappedNum)
		}
		m[mappedNum] = append(m[mappedNum], num)
	}

	sort.Ints(keys)

	for _, key := range keys {
		res = append(res, m[key]...)
	}

	return res
}
