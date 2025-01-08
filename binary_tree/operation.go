package binarytree

func (root *TreeNode) PreOrderTraversal(result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Value)
	root.Left.PreOrderTraversal(result)
	root.Right.PreOrderTraversal(result)
}

func (root *TreeNode) InOrderTraversal(result *[]int) {
	if root == nil {
		return
	}

	root.Left.PreOrderTraversal(result)
	*result = append(*result, root.Value)
	root.Right.PreOrderTraversal(result)
}

func (root *TreeNode) PostOrderTraversal(result *[]int) {
	if root == nil {
		return
	}

	root.Left.PreOrderTraversal(result)
	root.Right.PreOrderTraversal(result)
	*result = append(*result, root.Value)
}
