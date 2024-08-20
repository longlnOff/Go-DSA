package main

import "errors"

type DynamicArray struct {
	Elements 	[]any
	Length		int
	Capacity 	int
}

const (
	defaultValueCapacity = 10
)

func (da DynamicArray) isEmpty() bool {
	return da.Length == 0
}

func (da DynamicArray) checkRange(index int) error {
	if index >= da.Length || index < 0 {
		return errors.New("index out of range")
	} else {
		return nil
	}
}

func (da *DynamicArray) set(index int, value any) error {
	err := da.checkRange(index)
	if err != nil {
		return err
	} else {
		da.Elements[index] = value
		return nil
	}
}

func (da *DynamicArray) resizeArray() {
	if da.Capacity == 0 {
		da.Capacity = defaultValueCapacity
	} else {
		da.Capacity = da.Capacity << 1
	}

	newDataElement := make([]any, da.Capacity)
	copy(newDataElement, da.Elements)
	da.Elements = newDataElement
}

func (da *DynamicArray) append(element any) {
	if da.Length == da.Capacity {
		da.resizeArray()
	}

	da.Elements[da.Length] = element
	da.Length++
}

func (da *DynamicArray) removeAt(index int) error {
	err := da.checkRange(index)
	if err != nil {
		return err
	}

	copy(da.Elements[index:], da.Elements[index+1:])
	da.Elements[da.Length-1] = nil
	da.Length--
	return nil
}

func (da DynamicArray) get(index int) (any, error) {
	err := da.checkRange(index)
	if err != nil {
		return nil, err
	} else {
		return da.Elements[index], nil
	}
}

func (da DynamicArray) getData() []any {
	return da.Elements[:da.Length]
}

func (da DynamicArray) indexOf(value any) (int, error) {
	for i, v := range(da.Elements) {
		if v == value {
			return i, nil
		}
	}

	return -1, errors.New("array doesn't contain object")
}

func (da *DynamicArray) removeObject(value any) error {
	i, err := da.indexOf(value)
	if err != nil {
		return err
	} else {
		da.removeAt(i)
		return nil
	}
}