package helper

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeFromSlice(nums []int) *ListNode {
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
