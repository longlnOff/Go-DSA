package doubly

import (
	"errors"
	"fmt"
)

type DoublyLinkedList[T any] struct {
	length int
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Data T
	Next *Node[T]
	Previous *Node[T]
}

func NewNode[T any](Data T) *Node[T] {
	return &Node[T]{Data: Data, Next: nil, Previous: nil}
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{length: 0, Head: nil, Tail: nil}
}

func (ll *DoublyLinkedList[T]) GetLength() int {
	return ll.length
}

