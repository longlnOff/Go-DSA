package arrayqueue

import (
	"errors"
)

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() *Queue[T]{
	return &Queue[T]{}
}

func (q *Queue[T]) EnQueue(data T) {
	q.elements = append(q.elements, data)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	data := q.elements[0]
	q.elements = q.elements[1:]
	return data, nil
}

func (q *Queue[T]) FrontQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	return q.elements[0], nil
}


func (q *Queue[T]) BackQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	length := q.LenQueue()
	return q.elements[length-1], nil
}


func (q *Queue[T]) LenQueue() int {
	return len(q.elements)
}

func (q *Queue[T]) IsEmptyQueue() bool {
	if q.LenQueue() == 0 {
		return true
	} else {
		return false
	}
}

