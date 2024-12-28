package queue

type Queue struct {
	items   []int
	maxSize int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		items:   []int{},
		maxSize: maxSize,
	}
}
