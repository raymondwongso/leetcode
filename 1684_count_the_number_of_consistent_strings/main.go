package main

import "fmt"

func main() {
	allowed := "ab"
	words := []string{
		"ad", "bd", "aaab", "baa", "badab",
	}

	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	set := map[rune]struct{}{}

	for _, r := range allowed {
		set[r] = struct{}{}
	}

	count := 0

Wordloop:
	for _, word := range words {
		for _, c := range word {
			if _, ok := set[c]; !ok {
				continue Wordloop
			}
		}
		count += 1
	}

	return count
}
