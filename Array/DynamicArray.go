package main

import "errors"

type DynamicArray struct {
	elements 	[]any
	length		int
	capacity 	int
}

func (da DynamicArray) isEmpty() bool {
	return da.length == 0
}

func (da DynamicArray) checkRange(index int) error {
	if index >= da.length || index < 0 {
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
		da.elements[index] = value
		return nil
	}
}

func (da *DynamicArray) resizeArray() {
	if da.capacity == 0 {
		da.capacity = da.length
	} else {
		da.capacity = da.capacity << 1
	}

	newDataElement := make([]any, da.capacity)
	copy(newDataElement, da.elements)
	da.elements = newDataElement
}

func (da *DynamicArray) append(element any) {
	if da.length == da.capacity {
		da.resizeArray()
	}

	da.elements[da.length] = element
	da.length++
}

func (da *DynamicArray) removeAt(index int) error {
	err := da.checkRange(index)
	if err != nil {
		return err
	}

	copy(da.elements[index:], da.elements[index+1:])
	da.elements[da.length-1] = nil
	da.length--
	return nil
}

func (da DynamicArray) get(index int) (any, error) {
	err := da.checkRange(index)
	if err != nil {
		return nil, err
	} else {
		return da.elements[index], nil
	}
}

func (da DynamicArray) getData() []any {
	return da.elements[:da.length]
}

func (da DynamicArray) indexOf(value any) (int, error) {
	for i, v := range(da.elements) {
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