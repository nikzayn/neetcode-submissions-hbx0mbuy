/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//BFS + level order traversal
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}

    if root == nil {
        return res
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        qLen := len(queue)
        level := []int{}

        for i := 0; i < qLen; i++ {
            node := queue[0] //Top
            queue = queue[1:] //Pop

            level = append(level, node.Val) //Appending the node val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        res = append(res, level)
    }

    return res
}
