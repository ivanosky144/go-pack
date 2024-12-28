package stack

type Stack struct {
	items   []int
	maxSize int
}

func NewStack(maxSize int) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
