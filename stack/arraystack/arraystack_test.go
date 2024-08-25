package arraystack

import (
	"errors"
	"reflect"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T) {
	newStack := NewStack[int]()
	t.Run("Stack With Array", func(t *testing.T) {

		newStack.Push(2)
		newStack.Push(3)

		t.Run("Stack Push", func(t *testing.T) {
			var stackElements []any
			for i := 0; i < 2; i++ {
				poppedElement, _ := newStack.Pop()
				stackElements = append(stackElements, poppedElement)
			}

			if !reflect.DeepEqual([]any{3, 2}, stackElements) {
				t.Errorf("Stack Push is not work we expected %v but got %v", []any{3, 2}, newStack)
			}

			newStack.Push(2)
			newStack.Push(3)
		})

		pop, _ := newStack.Pop()

		t.Run("Stack Pop", func(t *testing.T) {
			if newStack.Length() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})

		newStack.Push(2)
		newStack.Push(83)

		t.Run("Stack Peak", func(t *testing.T) {
			if val, _ := newStack.Peek(); val != 83 {
				t.Errorf("Stack Peek is not work we expected %v but got %v", 83, val)
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if newStack.Length() != 3 {
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, newStack.Length())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if newStack.IsEmpty() != nil {
				t.Errorf("Stack Empty is not work we expected %v but got %v", nil, newStack.IsEmpty())
			}

			newStack.Pop()
			newStack.Pop()
			newStack.Pop()

			if newStack.IsEmpty() == nil {
				t.Errorf("Stack Empty is not work we expected %v but got %v", errors.New("stack is empty"), newStack.IsEmpty())
			}
		})
	})
}