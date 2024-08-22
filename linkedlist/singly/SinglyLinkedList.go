package singly

import (
	"errors"
	"fmt"
)

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

func (ll *SinglyLinkedList[T]) DelAtBeg() (T, error) {
	var temp_val T
	if ll.length != 0 {
		temp_val = ll.Head.Data
		if ll.length == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else {
			temp := ll.Head
			ll.Head = temp.Next
			temp.Next = nil
		}
		ll.length--
		return temp_val, nil
	}
	err := errors.New("linked list is empty, can't delete anymore")
	return temp_val, err
}

func (ll *SinglyLinkedList[T]) DelAtEnd() (T, error) {
	var temp_val T
	if ll.length != 0 {
		temp_val = ll.Tail.Data
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
		return temp_val, nil
	}
	err := errors.New("linked list is empty, can't delete anymore")
	return temp_val, err
}

func (ll *SinglyLinkedList[T]) CheckRangeFromIndex(index int) error {
	if index >= ll.length || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func (ll *SinglyLinkedList[T]) Display() {
	fmt.Print(ll.Head.Data)
	for run := ll.Head.Next; run != nil; run = run.Next {
		fmt.Print(" --> ", run.Data)
	}
	fmt.Print("\n")
}

func (ll *SinglyLinkedList[T]) DelByPos(pos int) (T, error) {
	var tempt T
	if err := ll.CheckRangeFromIndex(pos); err != nil {
		return tempt, err
	}
	if pos == 0 {
		return ll.DelAtBeg()
	} else if pos == ll.length - 1 {
		return ll.DelAtEnd()
	}
	index := 0
	for run := ll.Head; run != nil; run = run.Next {
		if index + 1 == pos {
			tempt_node := run.Next
			tempt = tempt_node.Data
			run.Next = tempt_node.Next
			tempt_node.Next = nil
			tempt_node = nil
			break
		} else {
			index++
		}
	}
	ll.length--
	return tempt, nil
}

func (ll *SinglyLinkedList[T]) GetLength() int {
	return ll.length
}

func (ll *SinglyLinkedList[T]) Reverse() error {
	if ll.length == 0 {
		return errors.New("linked list is empty")
	} else if ll.length == 1 {
		return nil
	} else {
		var prev, next *Node[T]		// initial value of prev and next is nil
		run := ll.Head
		ll.Tail = ll.Head
		for run != nil {
			next = run.Next
			run.Next = prev
			prev = run
			run = next
		}
		ll.Head = prev

		return nil
	}
}

func (ll *SinglyLinkedList[T]) ReversePartition(left, right int) error {
		err := ll.CheckRangeFromIndex(left)
		if err != nil {
			return err
		}
		tmpNode := &Node[T]{}
		tmpNode.Next = ll.Head
		pre := tmpNode
		for i := 0; i < left-1; i++ {
			pre = pre.Next
		}
		cur := pre.Next
		for i := 0; i < right-left; i++ {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		ll.Head = tmpNode.Next
		return nil
	}