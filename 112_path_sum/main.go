package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return recursive(root, 0, targetSum)
}

func recursive(node *TreeNode, sum, targetSum int) bool {
	if node == nil {
		return false
	}

	sum += node.Val
	if sum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	if recursive(node.Left, sum, targetSum) {
		return true
	}

	if recursive(node.Right, sum, targetSum) {
		return true
	}

	return false
}
