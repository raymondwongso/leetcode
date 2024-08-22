package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}

// because string only contains lowercase english letter, there are only 26 distinct values.
// we set a tally arrays for s and p respectively, with lenght of 26, each for each characters.
// first we count the first len(p) characters. if we found 'a', we increase the tally for 'a' (index 0)
// let arrs and arrp as the tally arrays for s and p respectively, if arrs and arrp is equal, it is anagram.
// continue iterate for each character in s, calculating tally and comparing with arrp.
// Time complexity: O(n) where n is the lenght of s
// Space complexity: O(2*26) for arrs and arpp, O(n) for the result array. Ultimately O(n)
func findAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	arrs, arrp := [26]int{}, [26]int{}
	offset := 97

	for i := range p {
		arrp[int(p[i])-offset] += 1
		arrs[int(s[i])-offset] += 1
	}

	res := []int{}
	if arrs == arrp {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		arrs[int(s[i-len(p)])-offset] -= 1
		arrs[int(s[i])-offset] += 1
		if arrs == arrp {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}
