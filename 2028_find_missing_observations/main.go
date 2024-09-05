package main

import "fmt"

func main() {
	rolls := []int{3, 2, 4, 3}
	mean := 4
	int := 2
	fmt.Println(missingRolls(rolls, mean, int))
}

// Time complexity: O(m + n)
// Space complexity: O(n)
// m = len of rolls, n = len of res
func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, r := range rolls {
		sum += r
	}

	targetSum := (len(rolls) + n) * mean
	diff := targetSum - sum

	// impossible case e.g:
	// diff: 21, n: 2
	// at most we can get 12 (6 * 2)
	if diff/6 > n {
		return []int{}
	}

	if diff/6 == n && diff%6 > 0 {
		return []int{}
	}

	// impossible case e.g:
	// diff: 2, n: 3
	// at most we can only get 2 1 roll to reach target sum, still missing 1 roll
	if diff < n {
		return []int{}
	}

	x := diff / n
	res := make([]int, n)

	for i := range res {
		res[i] = x
	}

	remainder := diff % n
	if remainder > 0 {
		for i := range res {
			if remainder == 0 {
				break
			}

			if remainder > 6-res[i] {
				remainder -= 6 - res[i]
				res[i] = 6
			} else {
				res[i] += remainder
				remainder = 0
			}
		}
	}

	return res
}
