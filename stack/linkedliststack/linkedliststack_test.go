package linkedliststack

import (
	"testing"
)

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T) {
	newStack := NewStackLL[int]()

	newStack.Push(1)
	newStack.Push(2)

	t.Run("Stack Push", func(t *testing.T) {
		result := newStack.Show()
		expected := []any{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
			}
		}
	})

	t.Run("Stack isEmpty", func(t *testing.T) {
		if newStack.IsEmpty() != nil {
			t.Error("Stack isEmpty is returned nil but expected 'stack is empty'", newStack.IsEmpty())
		}

	})

	t.Run("Stack Length", func(t *testing.T) {
		if newStack.Length() != 2 {
			t.Error("Stack Length should be 2 but got", newStack.Length())
		}
	})

	newStack.Pop()
	pop, _ := newStack.Pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if pop != 1 {
			t.Error("Stack Pop should return 1 but is returned", pop)
		}

	})

	newStack.Push(52)
	newStack.Push(23)
	newStack.Push(99)
	t.Run("Stack Peek", func(t *testing.T) {
		if a, _ := newStack.Peek(); a != 99 {
			t.Error("Stack Peak should return 99 but got ", a)
		}
	})
}