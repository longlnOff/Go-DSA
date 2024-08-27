package linkedlistqueue

import (
	"reflect"
	"testing"
)

// TestQueueArray for testing Stack with Array
func TestQueueLinkedList(t *testing.T) {
	queue1 := NewQueue[int]()

	queue1.EnQueue(2)
	queue1.EnQueue(4)
	queue1.EnQueue(6)
	queue1.EnQueue(7)
	queue1.EnQueue(8)
	queue1.EnQueue(9)

	t.Run("Test EnQueue", func(t *testing.T) {
		expected := []any{2,4,6,7,8,9}
		result := []any{}
		for i := queue1.head; i != nil; i = i.next {
			result = append(result, i.data)
		}
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("EnQueue is not work, we expected %v but got %v", expected, result)
		}
	})

	t.Run("Test FrontQueue", func(t *testing.T){
		expected := 2
		result, _ := queue1.FrontQueue()
		if expected != result {
			t.Errorf("FrontQueue is not work, we expected %v but got %v", expected, result)
		}

	})

	t.Run("Test BackQueue", func(t *testing.T){
		expected := 9
		result, _ := queue1.BackQueue()
		if expected != result {
			t.Errorf("BackQueue is not work, we expected %v but got %v", expected, result)
		}

	})


	queue1.DeQueue()
	queue1.DeQueue()
	queue1.DeQueue()
	t.Run("Test DeQueue", func(t *testing.T) {
		expected := []any{7,8,9}
		result := []any{}
		for i := queue1.head; i != nil; i = i.next {
			result = append(result, i.data)
		}
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("DeQueue is not work, we expected %v but got %v", expected, result)
		}
	})

	t.Run("Test IsEmptyQueue (not empty)", func(t *testing.T) {
		expected := false
		result := queue1.IsEmptyQueue()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("IsEmptyQueue is not work, we expected %v but got %v", expected, result)
		}
	})


	queue1.DeQueue()
	queue1.DeQueue()
	queue1.DeQueue()
	t.Run("Test IsEmptyQueue (empty)", func(t *testing.T) {
		expected := true
		result := queue1.IsEmptyQueue()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("IsEmptyQueue is not work, we expected %v but got %v", expected, result)
		}
	})


































}