package arrayqueue

import (
	"reflect"
	"testing"
)

// TestQueueArray for testing Stack with Array
func TestQueueArray(t *testing.T) {
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
		for _, val := range(queue1.elements) {
			result = append(result, val)
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
		for _, val := range(queue1.elements) {
			result = append(result, val)
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


	t.Run("Test EnQueue", func(t *testing.T) {
		var newQueue Queue[int]
		newQueue.EnQueue(2)
		newQueue.EnQueue(3)
		newQueue.EnQueue(4)
		newQueue.EnQueue(45)
		f, _ := newQueue.FrontQueue()
		b, _ := newQueue.BackQueue()
		if  f != 2 && b != 45 {
			t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 45, f, b)
		}

	})

	t.Run("Test DeQueue", func(t *testing.T) {
		var newQueue Queue[int]
		newQueue.EnQueue(2)
		newQueue.EnQueue(3)
		newQueue.EnQueue(4)

		newQueue.DeQueue()
		a, _ := newQueue.DeQueue()
		if a != 3 {
			t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 3, a)
		}

		//fmt.Println(newQueue.show())
	})

	t.Run("Test Queue isEmpty", func(t *testing.T) {
		var newQueue Queue[int]
		if newQueue.IsEmptyQueue() != true {
			t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, newQueue.IsEmptyQueue())
		}

		newQueue.EnQueue(3)
		newQueue.EnQueue(4)

		if newQueue.IsEmptyQueue() != false {
			t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, newQueue.IsEmptyQueue())
		}
	})

	t.Run("Test Queue Length", func(t *testing.T) {
		var newQueue Queue[int]
		if newQueue.LenQueue() != 0 {
			t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, newQueue.LenQueue())
		}

		newQueue.EnQueue(3)
		newQueue.EnQueue(4)
		newQueue.DeQueue()
		newQueue.EnQueue(22)
		newQueue.EnQueue(99)
		newQueue.DeQueue()
		newQueue.DeQueue()

		if newQueue.LenQueue() != 1 {
			t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, newQueue.LenQueue())
		}

	})































}