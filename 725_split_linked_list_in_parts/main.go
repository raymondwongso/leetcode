package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	k := 3

	fmt.Println(splitListToParts(head, k))
}

// Time complexity: O(m) where m is len(head)
// Space complexity: O(k)
func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	cur := head

	lenarr := 0
	// O(m) where m is len(head)
	for cur != nil {
		lenarr += 1
		cur = cur.Next
	}

	// O(k)
	// inner loop, O(m) where k == 1
	// hence worst case, O(m)
	for i := 0; i < len(res); i++ {
		if head == nil {
			break
		}

		res[i] = head

		x := lenarr / k
		if i < lenarr%k {
			x += 1
		}

		for j := 1; j < x; j++ {
			head = head.Next
		}

		temp := head.Next
		head.Next = nil
		head = temp
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
