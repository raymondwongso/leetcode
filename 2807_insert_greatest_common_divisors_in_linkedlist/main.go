package main

import (
	"fmt"

	"github.com/raymondwongso/leetcode/helper"
)

func main() {
	headarr := []int{18, 6, 10}
	head := helper.ListNodeFromSlice(headarr)
	res := insertGreatestCommonDivisors(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func insertGreatestCommonDivisors(head *helper.ListNode) *helper.ListNode {
	if head.Next == nil {
		return head
	}

	res := head
	prev := head
	head = head.Next

	for head != nil {
		node := &helper.ListNode{
			Val:  gcd(prev.Val, head.Val),
			Next: head,
		}
		prev.Next = node
		prev = head
		head = head.Next
	}

	return res
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}

	res := a % b
	if res == 0 {
		return b
	}

	return gcd(b, res)
}
