package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m, n := 3, 5
	nums := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
	head := createListNode(nums)

	res := spiralMatrix(m, n, head)
	for i := range res {
		for j := range res[i] {
			fmt.Print(res[i][j])
		}
		fmt.Println()
	}
}

// Time complexity: O(m*n)
// Space complexity: O(m*n)
func spiralMatrix(m, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	x, y, face, nextX, nextY := 0, -1, 0, 0, 0
	for head != nil {
		nextX, nextY = x+dirs[face][0], y+dirs[face][1]

		if nextX < 0 || nextX >= m {
			face = (face + 1) % 4
			continue
		}

		if nextY < 0 || nextY >= n {
			face = (face + 1) % 4
			continue
		}

		if res[nextX][nextY] != -1 {
			face = (face + 1) % 4
			continue
		}

		res[nextX][nextY] = head.Val
		x, y = nextX, nextY

		head = head.Next
	}

	return res
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var (
		head *ListNode
		prev *ListNode
	)
	for _, n := range nums {
		if prev == nil {
			head = &ListNode{
				Val: n,
			}
			prev = head
			continue
		}

		prev.Next = &ListNode{
			Val: n,
		}
		prev = prev.Next
	}

	return head
}
