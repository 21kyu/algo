func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || val == root.Val {
        return root
    } else if val > root.Val {
        return searchBST(root.Right, val)
    } else {
        return searchBST(root.Left, val)
    }
}