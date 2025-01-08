package stack

import "errors"

func (s *Stack) Push(item int) (*Stack, error) {
	if len(s.items) >= s.maxSize {
		return s, errors.New("Stack is full")
	}
	s.items = append(s.items, item)
	return s, nil
}

func (s *Stack) Pop() (*Stack, error) {
	if len(s.items) == 0 {
		return s, errors.New("Stack is empty")
	}
	s.items = s.items[1:]
	return s, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
