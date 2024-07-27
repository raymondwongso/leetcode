// // https://leetcode.com/problems/sort-an-array
// package main

// // Driver code
// func main() {
// 	// fmt.Println("https://leetcode.com/problems/sort-an-array")

// 	// // change input for your testcases.
// 	// input := []int{38, 27, 43, 3, 9, 82, 10}
// 	// fmt.Printf("Input: %+v\n", input)
// 	// fmt.Printf("Sorted array: %+v\n", MergeSort(input))
// }

// func MergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2

// 	// Copy arr, and slice it.
// 	// this way, even if arr modified, leftArr and rightArr stay the same.
// 	tempArr := make([]int, len(arr))
// 	copy(tempArr, arr)

// 	leftArr := tempArr[:mid]
// 	rightArr := tempArr[mid:]

// 	leftArr = MergeSort(leftArr)
// 	rightArr = MergeSort(rightArr)

// 	i, j, k := 0, 0, 0

// 	for i < len(leftArr) && j < len(rightArr) {
// 		if leftArr[i] < rightArr[j] {
// 			arr[k] = leftArr[i]
// 			i++
// 		} else {
// 			arr[k] = rightArr[j]
// 			j++
// 		}
// 		k++
// 	}

// 	for i < len(leftArr) {
// 		arr[k] = leftArr[i]
// 		i++
// 		k++
// 	}

// 	for j < len(rightArr) {
// 		arr[k] = rightArr[j]
// 		j++
// 		k++
// 	}

// 	return arr
// }

// // Bubble Sort	O(n^2) / O(n^2)	O(1)	Yes	Simple to implement, good for small datasets	Very slow for large datasets
// // Insertion Sort	O(n^2) / O(n^2)	O(1)	Yes	Simple, efficient for small or nearly sorted datasets	Inefficient for large datasets
// // Selection Sort	O(n^2) / O(n^2)	O(1)	No	Simple to understand, performs well on small datasets	Inefficient for large datasets
// // Merge Sort	O(n log n) / O(n log n)	O(n)	Yes	Consistent O(n log n) time, good for linked lists	Requires additional space, not in-place
// // Quick Sort	O(n log n) / O(n^2)	O(log n)	No	Fast average case, good cache performance	Worst-case is O(n^2), unstable
// // Heap Sort	O(n log n) / O(n log n)	O(1)	No	Efficient, good for large datasets, in-place	Unstable, slightly slower than Quick Sort on average
// // Counting Sort	O(n + k) / O(n + k)	O(k)	Yes	Very fast for small range of integers	Limited to integer keys, uses extra space
// // Radix Sort	O(nk) / O(nk)	O(n + k)	Yes	Linear time for small ranges, stable	Requires additional space, limited to integer keys
// // Bucket Sort	O(n + k) / O(n^2)	O(n)	Yes	Very efficient for uniformly distributed data	Performance degrades with non-uniform data
// // Shell Sort	O(n log n) to O(n^2) / O(n^2)	O(1)	No	More efficient than Insertion Sort, simple	Complex to choose gap sequence, unstable
package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	heapSort(arr)
	fmt.Printf("%+v\n", arr)
}
