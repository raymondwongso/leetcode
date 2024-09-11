package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(triangularSum(nums))
}

func triangularSum(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		for i := start; i < end; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		end -= 1
	}

	return nums[0]
}
