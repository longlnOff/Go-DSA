package singly

import "errors"

type SinglyLinkedList[T any] struct {
	length int
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func NewNode[T any](Data T) *Node[T] {
	return &Node[T]{Data, nil}
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{0, nil, nil}
}

func (ll *SinglyLinkedList[T]) AddAtBeg(Data T) error {
	node := NewNode(Data)
	if ll.length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
	ll.length++
	return nil
}

func (ll *SinglyLinkedList[T]) AddAtEnd(Data T) error {
	node := NewNode(Data)
	if ll.length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}
	ll.length++

	return nil
}

func (ll *SinglyLinkedList[T]) DelAtBeg() error {
	if ll.length != 0 {
		if ll.length == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else {
			temp := ll.Head
			ll.Head = temp.Next
			temp.Next = nil
		}
		ll.length--
		return nil
	}
	err := errors.New("linked list is empty, can't delete anymore")
	return err
}

func (ll *SinglyLinkedList[T]) DelAtEnd() error {
	if ll.length != 0 {
		if ll.length == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else {
			run := ll.Head
			for {
				if run.Next == ll.Tail {
					run.Next = nil
					ll.Tail = run
					break
				}
				run = run.Next
			}
		}
		ll.length--
		return nil
	}
	err := errors.New("linked list is empty, can't delete anymore")
	return err
}