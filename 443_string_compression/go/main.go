package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("https://leetcode.com/problems/string-compression/")

	// change input for your testcases.
	input := []byte{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}
	fmt.Printf("Input: %+v\n", input)
	fmt.Printf("Compress Result Length: %+v\n", StringCompress(input))
}

func StringCompress(chars []byte) int {
	n := len(chars)

	if n == 1 {
		return 1
	}

	i, j := 0, 0

	for j < n {
		count := 1
		for j < n-1 && chars[j] == chars[j+1] {
			count++
			j++
		}

		chars[i] = chars[j]
		i++

		if count > 1 {
			strCount := strconv.Itoa(count)
			for _, c := range strCount {
				chars[i] = byte(c)
				i++
			}
		}

		j++
	}

	str := ""
	for _, c := range chars {
		str = str + string(c)
	}
	fmt.Printf("chars result: %+v\n", str)
	return i
}
