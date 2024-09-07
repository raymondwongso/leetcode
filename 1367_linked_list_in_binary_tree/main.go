package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(isSubPath(head, root))
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	arrhead := []int{}
	for head != nil {
		arrhead = append(arrhead, head.Val)
		head = head.Next
	}

	arr := []*TreeNode{root}
	for len(arr) > 0 {
		node := arr[0]
		arr = arr[1:]

		if recursive(arrhead, 0, node) {
			return true
		}
		if node.Left != nil {
			arr = append(arr, node.Left)
		}

		if node.Right != nil {
			arr = append(arr, node.Right)
		}
	}

	return false
}

func recursive(arrhead []int, iarrhead int, node *TreeNode) bool {
	if node == nil {
		return false
	}

	if arrhead[iarrhead] == node.Val {
		if iarrhead+1 == len(arrhead) {
			return true
		}
		if recursive(arrhead, iarrhead+1, node.Left) {
			return true
		}
		if recursive(arrhead, iarrhead+1, node.Right) {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
