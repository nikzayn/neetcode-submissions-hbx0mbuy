/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	left := depth(root.Left)
	right := depth(root.Right)

	return abs(left - right) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left)
	right := depth(root.Right)

	return 1 + max(left, right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
