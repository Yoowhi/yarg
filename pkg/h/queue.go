package h

import "errors"

type Queue[T any] struct {
	list   []T
	Length int
}

func (q *Queue[T]) Enqueue(item T) {
	q.list = append(q.list, item)
	q.Length++
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("nothing to dequeue")
	}
	el := q.list[0]
	q.list = q.list[1:len(q.list)]
	q.Length--
	return &el, nil
}

func ConvertToQueue[T any](list []T) Queue[T] {
	return Queue[T]{list: list}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.list) == 0
}
