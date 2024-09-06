package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	res := modifiedList(nums, head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	set := map[int]struct{}{}
	for i := range nums {
		set[nums[i]] = struct{}{}
	}

	var (
		root    *ListNode
		prev    *ListNode
		current = head
	)

	for current != nil {
		if _, ok := set[current.Val]; ok {
			if prev != nil {
				prev.Next = nil
			}
			current = current.Next
			continue
		}

		if root == nil {
			root = current
		} else {
			prev.Next = current
		}

		prev = current
		current = current.Next
	}

	return root
}
