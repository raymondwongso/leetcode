package main

import "fmt"

func main() {
	chalk := []int{3, 4, 1, 2}
	k := 25
	fmt.Println(chalkReplacer(chalk, k))
}

// Get the sum of all chalk.
// if k is larger than sum, modulo k with sum so we get the remainder (put to k), the last leg of iteration
// for each chalk, check whether k is larger than current chalk or not.
// if larger, substract it and move on to the next chalk
// if not, that is the answer.
// Time complexity: O(n)
// Space complexity: O(1)
func chalkReplacer(chalk []int, k int) int {
	var sum int = 0

	for _, c := range chalk {
		sum += c
	}

	if k >= sum {
		k = k % sum
	}

	for i, c := range chalk {
		if k >= c {
			k -= c
		} else {
			return i
		}
	}

	return 0
}
