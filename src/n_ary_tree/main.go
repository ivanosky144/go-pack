package narytree

type Node struct {
	Value    int
	Children []*Node
}

type TreeNode struct {
	Root *Node
}

func NewNaryTree(value int) *TreeNode {
	return &TreeNode{
		Root: &Node{
			Value:    value,
			Children: []*Node{},
		},
	}
}
