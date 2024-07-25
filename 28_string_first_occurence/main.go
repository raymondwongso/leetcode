package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func main() {
	haystack := "mississippi"
	needle := "issip"

	fmt.Printf("Haystack: %+v\n", haystack)
	fmt.Printf("Needle: %+v\n", needle)
	fmt.Printf("First occurence index: %+v\n", stringFirstOccurence(haystack, needle))
}

func stringFirstOccurence(haystack string, needle string) int {
	i, j := 0, 0
	hLen, nLen := len(haystack), len(needle)

	if nLen > hLen {
		return -1
	}

	index := -1
	for i < hLen {
		fmt.Printf("i: %d , j: %d\n", i, j)
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i
			}
			j++
			fmt.Printf("same: index: %d | j: %d\n", index, j)
		} else {
			if index != -1 {
				i = index
			}

			index = -1
			j = 0
			fmt.Printf("not same: index: %d | j: %d\n", index, j)
		}

		if j == nLen {
			fmt.Printf("needle exhausted: index: %d | j: %d\n", index, j)
			return index
		}

		i++
	}

	return -1
}
