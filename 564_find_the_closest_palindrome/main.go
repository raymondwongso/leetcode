package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "11011"
	fmt.Println(nearestPalindromic(n))
}

func nearestPalindromic(s string) string {
	n, _ := strconv.Atoi(s)

	s1 := strconv.Itoa(n + 1)
	s2 := strconv.Itoa(n - 1)

	arr1 := make([]int, len(s1))
	for i := range s1 {
		arr1[i], _ = strconv.Atoi(string(s1[i]))
	}

	arr2 := make([]int, len(s2))
	for i := range s2 {
		arr2[i], _ = strconv.Atoi(string(s2[i]))
	}

	for i, j := 0, len(arr1)-1; i < j; i, j = i+1, j-1 {
		if arr1[j] <= arr1[i] {
			arr1[j] = arr1[i]
		} else {
			ri := 1
			for {
				arr1[j-ri] += 1
				if arr1[j-ri] > 9 {
					arr1[j-ri] = 0
					ri += 1
				} else {
					break
				}
			}

			arr1[j] = arr1[i]
		}
	}

	for i, j := 0, len(arr2)-1; i < j; i, j = i+1, j-1 {
		if arr2[j] >= arr2[i] {
			arr2[j] = arr2[i]
		} else {
			ri := 1
			for {
				arr2[j-ri] -= 1
				if arr1[j-ri] < 0 {
					arr1[j-ri] = 0
					ri += 1
				} else {
					break
				}
			}

			if len(arr2) == 2 && arr2[j-ri] == 0 {
				return "9"
			}
			arr2[j] = arr2[i]
		}
	}

	narr1 := 0
	for _, x := range arr1 {
		narr1 = narr1*10 + x
	}

	narr2 := 0
	for _, x := range arr2 {
		narr2 = narr2*10 + x
	}

	x := n - narr1
	if x < 0 {
		x *= -1
	}

	y := n - narr2
	if y < 0 {
		y *= -1
	}

	if y <= x {
		return strconv.Itoa(narr2)
	}

	return strconv.Itoa(narr1)
}
