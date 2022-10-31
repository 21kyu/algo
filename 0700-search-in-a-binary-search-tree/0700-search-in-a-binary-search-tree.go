func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    } else if root.Val < val {
        return searchBST(root.Right, val)
    } else {
        return searchBST(root.Left, val)
    }
}