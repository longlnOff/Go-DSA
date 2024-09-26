package main

import "golang.org/x/exp/constraints"

// BST data structure
type BSTNode[T constraints.Ordered] struct {
	Data T
	Left *BSTNode[T]
	Right *BSTNode[T]
}

func NewNodeBST[T constraints.Ordered](Data T) *BSTNode[T] {
	return &BSTNode[T]{Data: Data}
}

type BST[T constraints.Ordered] struct {
	Root *BSTNode[T]
}


func (bst *BST[T]) Insert(Data T) {
	if bst.Root == nil {
		bst.Root = NewNodeBST(Data)
	} else {
		parent := bst.Root
		current := bst.Root
		for current != nil {
			parent = current
			if Data < current.Data {
				current = current.Left
			} else {
				current = current.Right
			}
		}
		if Data < parent.Data {
			parent.Left = NewNodeBST(Data)
		} else {
			parent.Right = NewNodeBST(Data)
		}
	}
}

