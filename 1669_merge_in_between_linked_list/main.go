package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Submission result: 86ms (92.44%) | 7.11MB (87.40%)
// Time complexity: O(m+n)
// Space complexity: O(1)
func main() {
	// list1 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}
	// list2 := &ListNode{Val: 100, Next: &ListNode{Val: 101, Next: &ListNode{Val: 102, Next: &ListNode{Val: 103}}}}
	// res := mergeInBetween(list1, 2, 5, list2)

	list1 := &ListNode{Val: 10, Next: &ListNode{Val: 1, Next: &ListNode{Val: 13, Next: &ListNode{Val: 6, Next: &ListNode{Val: 9, Next: &ListNode{Val: 5}}}}}}
	list2 := &ListNode{Val: 100, Next: &ListNode{Val: 101, Next: &ListNode{Val: 102}}}
	res := mergeInBetween(list1, 3, 4, list2)
	printListNode(res)
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	i := 0
	rootList1 := list1
	for true {
		if i == a-1 {
			diff := b - a
			counter := 0
			i += 1
			runningList := list1.Next
			var tail *ListNode

			for true {
				if counter < diff {
					i += 1
					runningList = runningList.Next
					counter += 1
				} else {
					if runningList.Next != nil {
						tail = runningList.Next
					}
					break
				}
			}

			rootList2 := list2
			for list2.Next != nil {
				list2 = list2.Next
			}
			list2.Next = tail
			list1.Next = rootList2
			break
		}

		i += 1
		list1 = list1.Next
	}

	return rootList1
}

func printListNode(list *ListNode) {
	if list == nil {
		return
	}

	fmt.Println(list.Val)
	if list.Next != nil {
		printListNode(list.Next)
	}
}
