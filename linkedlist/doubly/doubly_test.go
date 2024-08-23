package doubly

import (
	"reflect"
	"testing"
)


func TestMain(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	// Test AddAtBeg
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)
	list.AddAtBeg(4)
	list.AddAtBeg(5)
	t.Run("Test AddAtBeg()", func(t *testing.T) {
		want := []any{5, 4, 3, 2, 1}
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

	// Test AddAtEnd
	list.AddAtEnd(1)
	list.AddAtEnd(2)
	list.AddAtEnd(3)
	list.AddAtEnd(4)
	list.AddAtEnd(5)
	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []any{5, 4, 3, 2, 1, 1, 2, 3, 4, 5}
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

	// Test DelAtBeg
	list.DelAtBeg()
	list.DelAtBeg()
	list.DelAtBeg()
	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := []any{2, 1, 1, 2, 3, 4, 5}
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

	// Test DelAtEnd
	list.DelAtEnd()
	list.DelAtEnd()
	list.DelAtEnd()
	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := []any{2, 1, 1, 2}
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

	// Test DelByPos
	list.DelByPos(1)
	list.DelByPos(0)
	list.DelByPos(1)
	t.Run("Test DelByPos()", func(t *testing.T) {
		want := []any{1}
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


	// Test Reverse
	list1 := NewDoublyLinkedList[int]()
	// Test AddAtBeg
	list1.AddAtBeg(1)
	list1.AddAtBeg(2)
	list1.AddAtBeg(3)
	list1.AddAtBeg(4)
	list1.AddAtBeg(5)
	list1.Reverse()
	t.Run("Test Reverse()", func(t *testing.T) {
		// from head to Tail
		want := []any{1, 2, 3, 4, 5}
		got := []any{}
		current := list1.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
		// from Tail to head
		want_tail := []any{5, 4, 3, 2, 1}
		got_tail := []any{}
		current_tail := list1.Tail
		got_tail = append(got_tail, current_tail.Data)
		for current.Previous != nil {
			current = current.Previous
			got_tail = append(got_tail, current.Data)
		}
		if !reflect.DeepEqual(got_tail, want_tail) {
			t.Errorf("got: %v, want: %v", got_tail, want_tail)
		}
	})

		// Test Reverse
		list2 := NewDoublyLinkedList[int]()
		// Test AddAtBeg
		list2.AddAtBeg(1)
		list2.AddAtBeg(2)
		list2.AddAtBeg(3)
		list2.AddAtBeg(4)
		list2.AddAtBeg(5)
		list2.AddAtBeg(6)
		list2.AddAtBeg(7)
		list2.AddAtBeg(8)
		list2.ReversePartition(2, 5)
		list2.Display()
		list2.DisplayReverse()
		t.Run("Test ReversePartition()", func(t *testing.T) {
			// from head to Tail
			want := []any{8, 7, 3, 4, 5, 6, 2, 1}
			got := []any{}
			current := list2.Head
			got = append(got, current.Data)
			for current.Next != nil {
				current = current.Next
				got = append(got, current.Data)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// from Tail to head
			want_tail := []any{1, 2, 6, 5, 4, 3, 7, 8}
			got_tail := []any{}
			current_tail := list2.Tail
			got_tail = append(got_tail, current_tail.Data)
			for current.Previous != nil {
				current = current.Previous
				got_tail = append(got_tail, current.Data)
			}
			if !reflect.DeepEqual(got_tail, want_tail) {
				t.Errorf("got: %v, want: %v", got_tail, want_tail)
			}
		})

}