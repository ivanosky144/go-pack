package narytree

func (n *Node) AddChild(value int) {
	child := &Node{
		Value:    value,
		Children: []*Node{},
	}
	n.Children = append(n.Children, child)
}
