package doubly

import (
	"errors"
	"fmt"
	// "fmt"
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

func (ll *DoublyLinkedList[T]) AddAtBeg(data T) error {
	newNode := NewNode(data)
	if ll.length == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		newNode.Previous = ll.Head.Previous
		ll.Head.Previous = newNode
		ll.Head = newNode
	}
	ll.length++

	return nil
}

func (ll *DoublyLinkedList[T]) AddAtEnd(data T) error {
	newNode := NewNode(data)
	if ll.length == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Tail.Next
		newNode.Previous = ll.Tail
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.length++

	return nil
}

func (ll *DoublyLinkedList[T]) DelAtBeg() (T, error) {
	var retData T
	if ll.length == 0 {
		return retData, errors.New("linked list is empty, can not delete")
	} else {
		retData = ll.Head.Data
		nextHead := ll.Head.Next
		nextHead.Previous = ll.Head.Previous
		ll.Head.Next = nil
		ll.Head = nextHead
	}
	ll.length--

	return retData, nil
}

func (ll *DoublyLinkedList[T]) DelAtEnd() (T, error) {
	var retData T
	if ll.length == 0 {
		return retData, errors.New("linked list is empty, can not delete")
	} else {
		retData = ll.Tail.Data
		prevEnd := ll.Tail.Previous
		prevEnd.Next = ll.Tail.Next
		ll.Tail.Previous = nil
		ll.Tail = prevEnd
	}
	ll.length--

	return retData, nil
}

func (ll *DoublyLinkedList[T]) DelByPos(pos int) (T, error) {
	var retData T
	if pos < 0 || pos >= ll.length {
		return retData, errors.New("invalid delete position")
	} else if pos == 0 {
		return ll.DelAtBeg()
	} else if pos == ll.length - 1 {
		return ll.DelAtEnd()
	} else {
		tral := ll.Head
		for countPos := 0; countPos <= pos; countPos++ {
			tral = tral.Next
		}
		retData = tral.Data
		prev := tral.Previous
		next := tral.Next
		tral.Previous = nil
		tral.Next = nil
		prev.Next = next
		next.Previous = prev
		ll.length--

		return retData, nil
	}
}

func (ll *DoublyLinkedList[T]) Reverse() error {
	if ll.length == 0 {
		return errors.New("linked list is empty, can not delete")
	} else if ll.length == 1 {
		return nil
	} else {
		var prev, next *Node[T]		// initial value of prev and next is nil
		cur := ll.Head
		ll.Tail = ll.Head
		for cur != nil {
			next = cur.Next
			cur.Next = prev
			cur.Previous = next
			prev = cur
			cur = next
		}
		ll.Head = prev

		return nil
	}
}

func (ll *DoublyLinkedList[T]) ReversePartition(left int, right int) error {
	if ll.length == 0 {
		return errors.New("linked list is empty, can not delete")
	} else if ll.length == 1 {
		return errors.New("linked list's length is one")
	}
	if err := ll.CheckRangeFromIndex(left, right); err != nil {
		return err
	}

	cur := ll.Head
	var prev *Node[T]
	var next *Node[T]
	for i := 0; i < left; i++ {
		cur = cur.Next
	}
	prev = cur.Previous
	for i := left; i < right; i++ {
		next = cur.Next
		cur.Next = next.Next
		next.Next.Previous = next.Previous
		next.Next = prev.Next
		prev.Next.Previous = next
		next.Previous = prev
		prev.Next = next
		
	}

	return nil
}


func (ll *DoublyLinkedList[T]) CheckRangeFromIndex(left int, right int) error {
	if left < 0 || right >= ll.length {
		return fmt.Errorf("invalid index: left: %v, right: %v", left, right)
	} else if left >= right {
		return fmt.Errorf("left value (%v) must less than right value (%v)", left, right)
	} else {
		return nil
	}
}

func (ll *DoublyLinkedList[T]) Display() {
	fmt.Println("Display")
	fmt.Print("Head ")
	for i := ll.Head; i != nil; i = i.Next {
		fmt.Print(i.Data, " ")
	}
	fmt.Println("Tail")
}

func (ll *DoublyLinkedList[T]) DisplayReverse() {
	fmt.Println("Display Reverse")
	fmt.Print("Tail ")
	for i := ll.Tail; i != nil; i = i.Previous {
		fmt.Print(i.Data, " ")
	}
	fmt.Println("Head")
}