package main

import "fmt"

type Solution struct {
	board       [][]byte
	wordMap     map[byte]bool
	tileVisited map[uint16]bool
	word        string
}

// https://leetcode.com/problems/word-search/
func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCFB"

	sol := &Solution{
		board:       board,
		word:        word,
		wordMap:     wordMap(word),
		tileVisited: map[uint16]bool{},
	}

	fmt.Printf("%t\n", sol.exist(board, word))
}

func (s *Solution) exist(board [][]byte, _ string) bool {
	m, n := len(board), len(board[0])
	wi := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s.backtrack(i, j, wi) {
				return true
			}
		}
	}

	return false
}

func (s *Solution) backtrack(i, j, wi int) bool {
	res := false
	tileVisitedKey := uint16(i)<<8 | uint16(j)
	if _, ok := s.wordMap[s.board[i][j]]; !ok {
		return false
	}

	visited, ok := s.tileVisited[tileVisitedKey]
	if ok {
		if visited {
			return false
		}
	}

	if s.board[i][j] == s.word[wi] {
		s.tileVisited[tileVisitedKey] = true
		if wi == len(s.word)-1 {
			return true
		}

		// top
		if i > 0 {
			res = s.backtrack(i-1, j, wi+1)
			if res {
				return res
			}
		}

		// right
		if j < len(s.board[0])-1 {
			res = s.backtrack(i, j+1, wi+1)
			if res {
				return res
			}
		}

		// bottom
		if i < len(s.board)-1 {
			res = s.backtrack(i+1, j, wi+1)
			if res {
				return res
			}
		}

		// left
		if j > 0 {
			res = s.backtrack(i, j-1, wi+1)
			if res {
				return res
			}
		}
	}

	s.tileVisited[tileVisitedKey] = false
	return res
}

func wordMap(word string) map[byte]bool {
	m := map[byte]bool{}

	for i := range word {
		m[word[i]] = true
	}

	return m
}
