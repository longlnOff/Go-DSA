package linkedliststack

import (
	"errors"
)

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func NewNode[T any](Data T) *Node[T] {
	return &Node[T]{Data, nil}
}

type StackLinkedList[T any] struct {
	length int
	head *Node[T]
}

func NewStackLL[T any] () *StackLinkedList[T] {
	return &StackLinkedList[T]{}
}

func (s *StackLinkedList[T]) Length() int {
	return s.length
}

func (s *StackLinkedList[T]) IsEmpty() error {
	if s.Length() == 0 {
		return errors.New("stack is empty")
	} else {
		return nil
	}
}


func (s *StackLinkedList[T]) Push(val T) {
	newNode := NewNode(val)
	defer func() {
		s.length++
	}()
	if s.IsEmpty() != nil {
		s.head = newNode
	} else {
		newNode.Next = s.head
		s.head = newNode
	}
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	var zeroVal T
	if err := s.IsEmpty(); err != nil {
		return zeroVal, err
	} else {
		zeroVal = s.head.Data
		temp := s.head
		s.head = temp.Next
		temp.Next = nil

		return zeroVal, nil
	}
}

func (s *StackLinkedList[T]) Peek() (T, error) {
	if err := s.IsEmpty(); err != nil {
		var zeroVal T
		return zeroVal, err
	} else {
		return s.head.Data, nil
	}
}

func (s *StackLinkedList[T]) Show() []T {
	var retSlice []T
	for i := s.head; i != nil; i = i.Next {
		retSlice = append(retSlice, i.Data)
	}

	return retSlice
}