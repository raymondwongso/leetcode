package main

import "fmt"

func main() {
	s := "ellllllllllllllellll"
	fmt.Println(findTheLongestSubstring(s))
}

// We store bitmap for each vowel
// e.g: 11111 for o, e, u, i, a
// if bitmask is on (value is 1), it means the vowel has shown odd-times
// if bitmask is off (value is 0), it means the vowel has shown even-times
// we store the first ever bitmask index found to m array of 32 int, 32 is the maximum bitmask value (11111)
// if we found vowel, we flip the bit for correspondence character. e.g: a, flip the 1 bit. i flip the 2 bit. u flip the 4 bit.
// for each bitmask found, check whether the bitmask already stored in m. This can be checked whether m[bitmask] value is "zero" or not
//
//	(in this case, -2 to differentiate between normal index value)
//
// if not found, store, this will be the earliest bitmask index found.
// iterate string. whenever we found specific bitmask again, we can be sure that the currentIndex - bitmap index stored earliest is the longest string
// time complexity: O(n) where n is len(s)
// space complexity: O(32) or O(1)
func findTheLongestSubstring(s string) int {
	m := [32]int{}
	for i := range m {
		m[i] = -2
	}
	m[0] = -1

	res, bitmask := 0, 0
	for i := range s {
		switch s[i] {
		case 'a':
			bitmask ^= 1
		case 'i':
			bitmask ^= 2
		case 'u':
			bitmask ^= 4
		case 'e':
			bitmask ^= 8
		case 'o':
			bitmask ^= 16
		}

		if m[bitmask] == -2 {
			m[bitmask] = i
		} else {
			if i-m[bitmask] > res {
				res = i - m[bitmask]
			}
		}
	}

	return res
}
