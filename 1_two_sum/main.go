package main

import "fmt"

// https://leetcode.com/problems/two-sum/
func main() {
	input := []int{3, 2, 4}
	target := 6
	fmt.Printf("%+v\n", hashMapSolution(input, target))
}

// Naive Solution
//
// For each nums[i] where i is 0 to len of nums,
// iterate nums[j] where j is i+1 to len of nums.
// if found any matching sum between nums[i] and nums[j] that make the target, return i and j
// Time complexity:  O(n2)
// Space complexity: O(1)
//
//	excluding input size, additional space required is
//	for i, n, j, n2 which is constant
// Submission result: 30ms (26.77%) | 3.58MB (92.43%)
func naiveSolution(nums []int, target int) []int {
	lenNums := len(nums)
	for i, n := range nums {
		for j := i + 1; j < lenNums; j++ {
			n2 := nums[j]
			if n+n2 == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Improved Hashmap Solution
//
// While populating hashmap, also lookup
// Time complexity: O(n)
// Space complexity: O(n)
// Submission result: 6ms (76.07%) | 4.25MB (19.92%)
func hashMapSolution(nums []int, target int) []int {
	hashMap := map[int]int{}

	for i, n := range nums {
		lookup := target - n
		if res, ok := hashMap[lookup]; ok {
			return []int{i, res}
		}

		hashMap[n] = i
	}

	return []int{}
}
