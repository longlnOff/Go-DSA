package singly

import (
	"fmt"
	"reflect"
	"testing"
)


func TestMain(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)

	t.Run("Test AddAtBeg()", func(t *testing.T) {
		want := []any{3, 2, 1}
		got := []any{}
		current := list.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.AddAtEnd(4)

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []any{3, 2, 1, 4}
		got := []any{}
		current := list.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := any(3)
		got, ok := list.DelAtBeg()
		if ok != nil {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := any(4)
		got, ok := list.DelAtEnd()
		if ok != nil {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Display()", func(t *testing.T) {
		want := 0
		got := 0
		list.Display()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test GetLength()", func(t *testing.T) {
		want := 2
		got := list.GetLength()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list2 := NewSinglyLinkedList[int]()
	list2.AddAtBeg(1)
	list2.AddAtBeg(2)
	list2.AddAtBeg(3)
	list2.AddAtBeg(4)
	list2.AddAtBeg(5)
	list2.AddAtBeg(6)
	list2.Display()
	got, err := list2.DelByPos(1)
	list2.Display()

	t.Run("Test DelAtPos()", func(t *testing.T) {
		want := 2
		if err != nil {
			t.Errorf("got: %v, want: %v", got, want)
		}
		want = 5
		got = list2.GetLength() 
		if err != nil {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("Test Reverse()", func(t *testing.T) {
		want := []any{1, 2, 3, 4, 6}
		err := list2.Reverse()
		if err != nil {
			t.Errorf("%v", err)
		}
		got := []any{}
		for run := list2.Head; run != nil; run = run.Next {
			got = append(got, run.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
		fmt.Println(list2.Tail.Data)
		fmt.Println(list2.Head.Data)
	})
	list3 := NewSinglyLinkedList[int]()
	list3.AddAtBeg(1)
	list3.AddAtBeg(2)
	list3.AddAtBeg(3)
	list3.AddAtBeg(4)
	list3.AddAtBeg(5)
	list3.AddAtBeg(6)
	list3.AddAtBeg(7)
	list3.AddAtBeg(8)
	list3.AddAtBeg(9)
	list3.AddAtBeg(10)
	
	list3.Display()
	t.Run("Test ReversePartition()", func(t *testing.T) {
		want := []any{10, 6, 7, 8, 9, 5, 4, 3, 2, 1}
		got := []any{}
		err := list3.ReversePartition(2, 5)

		if err != nil {
			t.Errorf("Incorrect boundary conditions entered%v", err)
		}
		current := list3.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}