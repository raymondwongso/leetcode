package main

import "fmt"

func main() {
	input := []int{3, 2, 4}
	target := 6
	fmt.Printf("%+v", hashMapSolution(input, target))
}

// Naive solution
// We iterate all possibility
// Time complexity:  O(n2)
// Space complexity: O(1)
//
//	excluding input size, additional space required is
//	for i, n, j, n2 which is constant
// Submission result: 30ms (26.77%) | 3.58MB (92.43%)
// func naiveSolution(nums []int, target int) []int {
// 	lenNums := len(nums)
// 	for i, n := range nums {
// 		for j := i + 1; j < lenNums; j++ {
// 			n2 := nums[j]
// 			if n+n2 == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{}
// }

// Improved Hashmap solution
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
