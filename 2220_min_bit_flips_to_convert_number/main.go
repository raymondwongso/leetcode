package main

import "fmt"

func main() {
	start := 10
	goal := 7
	fmt.Println(minBitFlips(start, goal))
}

func minBitFlips(start int, goal int) int {
	res := 0
	xor := start ^ goal
	for xor != 0 {
		res += xor & 1
		xor >>= 1
	}
	return res
}
