package linkedlistqueue

import (
	"errors"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T any] (data T) *Node[T] {
	return &Node[T]{data: data, next: nil}
}

type Queue[T any] struct {
	length int
	head *Node[T]
	tail *Node[T]
}

func NewQueue[T any]() *Queue[T]{
	return &Queue[T]{}
}

func (q *Queue[T]) EnQueue(data T) {
	defer func() {
		q.length++
	}()
	newNode := NewNode[T](data)
	if q.IsEmptyQueue() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	retVal := q.head.data
	q.head = q.head.next
	q.length--
	
	return retVal, nil
}

func (q *Queue[T]) FrontQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	return q.head.data, nil
}


func (q *Queue[T]) BackQueue() (T, error) {
	if q.IsEmptyQueue() {
		var zeroVal T
		return zeroVal, errors.New("queue is empty")
	}
	return q.tail.data, nil
}


func (q *Queue[T]) LenQueue() int {
	return q.length
}

func (q *Queue[T]) IsEmptyQueue() bool {
	if q.LenQueue() == 0 {
		return true
	} else {
		return false
	}
}

