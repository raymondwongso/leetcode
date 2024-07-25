package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring/
func main() {
	input := "bacab"
	fmt.Printf("%s\n", longestPalindromicSubstring(input))
}

func longestPalindromicSubstring(s string) string {
	ln := len(s)

	if ln == 1 {
		return s
	}

	arr := make([][]int, ln)
	for i := range arr {
		t := make([]int, ln)
		t[i] = 1
		arr[i] = t
	}

	// for _, row := range arr {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("=================")

	max := 1
	maxrow, maxcol := 0, 0
	step := 1

	for irow := 0; irow+step < ln; {
		// fmt.Printf("irow: %d | irow+step: %d | char: %s\n", irow, irow+step, s[irow:irow+step+1])
		if s[irow] == s[irow+step] {
			// for 2 length string
			if step == 1 {
				temp := 2 // 2 length string, 2 length palindrome
				arr[irow][irow+step] = temp
				if max < temp {
					max = temp
					maxrow = irow
					maxcol = irow + step
				}
			} else {
				// more than 2 length string, must check diagonal left first whether it is palindrom or not.
				// if diagonal left (subpalindrom in the middle) is palindrom, add the value with 2.
				// if not palindrom, then ignore (it will fallback to default max which is 1)
				dl := arr[irow+1][irow+step-1]
				if dl > 0 {
					temp := dl + 2
					arr[irow][irow+step] = temp
					if max < temp {
						max = temp
						maxrow = irow
						maxcol = irow + step
					}
				}
			}
		} else {
			// if first char and last char is not the same, then the palindrom is
			// s[0:i-1] or s[1:i] (the left and bottom from current position)
			// as we use `max, maxrow and maxcol` variable, we already cache the result.
			// we do not need to do anything for this case. The else clause is put only for documentation.
		}

		irow = irow + 1

		if irow+step >= ln {
			if step >= ln {
				break
			}

			irow = 0
			step = step + 1
		}
	}

	// this is for drawing the board
	// for _, row := range arr {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Printf("maxrow: %d | maxcol: %d", maxrow, maxcol)
	// fmt.Print(s[maxrow : maxcol+1])
	return s[maxrow : maxcol+1]
}
