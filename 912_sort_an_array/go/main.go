// https://leetcode.com/problems/sort-an-array
package main

import "fmt"

// Driver code
func main() {
	fmt.Println("https://leetcode.com/problems/sort-an-array")

	// change input for your testcases.
	input := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("Input: %+v\n", input)
	fmt.Printf("Sorted array: %+v\n", MergeSort(input))
}

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2

		// Copy arr, and slice it.
		// this way, even if arr modified, leftArr and rightArr stay the same.
		tempArr := make([]int, len(arr))
		copy(tempArr, arr)

		leftArr := tempArr[:mid]
		rightArr := tempArr[mid:]

		leftArr = MergeSort(leftArr)
		rightArr = MergeSort(rightArr)

		i, j, k := 0, 0, 0

		for i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] < rightArr[j] {
				arr[k] = leftArr[i]
				i++
			} else {
				arr[k] = rightArr[j]
				j++
			}
			k++
		}

		for i < len(leftArr) {
			arr[k] = leftArr[i]
			i++
			k++
		}

		for j < len(rightArr) {
			arr[k] = rightArr[j]
			j++
			k++
		}
	}

	return arr
}
