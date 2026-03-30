/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    calculateDiameter(root, &maxDiameter)
    return maxDiameter
}

func calculateDiameter(root *TreeNode, maxDiameter *int) int {
    if root == nil {
        return 0
    }

    left := calculateDiameter(root.Left, maxDiameter)
    right := calculateDiameter(root.Right, maxDiameter)

    if left+right > *maxDiameter {
        *maxDiameter = left+right
    }

    return 1 + max(left, right)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
