package singly

import (
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

	// t.Run("Test DelAtBeg()", func(t *testing.T) {
	// 	want := any(3)
	// 	got, ok := list.DelAtBeg()
	// 	if !ok {
	// 		t.Error("unexpected not-ok")
	// 	}
	// 	if got != want {
	// 		t.Errorf("got: %v, want: %v", got, want)
	// 	}
	// })

	// t.Run("Test DelAtEnd()", func(t *testing.T) {
	// 	want := any(4)
	// 	got, ok := list.DelAtEnd()
	// 	if !ok {
	// 		t.Error("unexpected not-ok")
	// 	}
	// 	if got != want {
	// 		t.Errorf("got: %v, want: %v", got, want)
	// 	}
	// })

	// t.Run("Test Count()", func(t *testing.T) {
	// 	want := 2
	// 	got := list.Count()
	// 	if got != want {
	// 		t.Errorf("got: %v, want: %v", got, want)
	// 	}
	// })

	// list2 := Singly[int]{}
	// list2.AddAtBeg(1)
	// list2.AddAtBeg(2)
	// list2.AddAtBeg(3)
	// list2.AddAtBeg(4)
	// list2.AddAtBeg(5)
	// list2.AddAtBeg(6)
}