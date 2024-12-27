package binarytree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
