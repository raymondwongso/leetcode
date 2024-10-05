package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("abc", "aadbcasext"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	offset := 97
	s1arr, s2arr := [26]int{}, [26]int{}

	for i := range s1 {
		s1arr[int(s1[i])-offset] += 1
		s2arr[int(s2[i])-offset] += 1
	}

	if s1arr == s2arr {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		j := i - len(s1)

		s2arr[int(s2[j])-offset] -= 1
		s2arr[int(s2[i])-offset] += 1

		if s1arr == s2arr {
			return true
		}
	}

	return false
}
