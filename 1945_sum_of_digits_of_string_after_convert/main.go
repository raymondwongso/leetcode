package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "leetcode"
	k := 2
	fmt.Println(getLucky(s, k))
}

func getLucky(s string, k int) int {
	var str string
	for i := range s {
		str += strconv.Itoa(int(s[i]) - 96)
	}

	fmt.Printf("%s\n", str)

	num := 0
	for _, r := range str {
		num += int(r - '0')
	}

	sum := 0
	for range k - 1 {
		for num > 0 {
			sum += (num % 10)
			num /= 10
		}
		num = sum
		sum = 0
	}

	return num
}
