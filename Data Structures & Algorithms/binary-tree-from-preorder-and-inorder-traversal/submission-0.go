/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    //preorder
    //Root -> left -> right
    //Inorder
    //Left -> Root -> Right
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    //create a root wit
    root := &TreeNode{Val: preorder[0]}

    mid := 0

    //find the mid in inorder
    for i, val := range inorder {
        if val == preorder[0] {
            mid = i
            break
        }
    }

    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

    return root
}
