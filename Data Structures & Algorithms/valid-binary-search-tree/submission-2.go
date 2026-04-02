/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt32, math.MaxInt32)
}

func isValid(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	} else if lower < node.Val && node.Val < upper {
		return isValid(node.Left, lower, node.Val) && isValid(node.Right, node.Val, upper)
	}
	return false
}
