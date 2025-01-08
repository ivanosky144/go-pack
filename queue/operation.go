package queue

import "errors"

func (q *Queue) Dequeue() (*Queue, error) {
	if len(q.items) == 0 {
		return q, errors.New("Queue is empty")
	}

	q.items = q.items[1:]
	return q, nil
}

func (q *Queue) Enqueue(value int) (*Queue, error) {
	if q.maxSize > 0 && len(q.items) >= q.maxSize {
		return q, errors.New("Queue is full")
	}
	q.items = append(q.items, value)
	return q, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("Queue is empty")
	}

	return q.items[0], nil
}
