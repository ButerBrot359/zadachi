var res int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res = 0

	recursive(root)

	return res
}

func recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l, r := recursive(node.Left), recursive(node.Right)
	res = max(res, l+r)

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}