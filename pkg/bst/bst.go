package main

import (
	"fmt"
)

// Node represents a single node in the BST
type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

// BST represents the binary search tree
type BST struct {
	root *BSTNode
}

// Insert inserts a new value into the BST
func (bst *BST) Insert(value int) {
	// if empty create new node
	if bst.root == nil {
		bst.root = &BSTNode{value: value}
	} else { // otherwise recurse
		bst.root.insert(value)
	}
}

// insert is a helper function to insert a value into the tree
func (n *BSTNode) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &BSTNode{value: value}
		} else {
			n.left.insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = &BSTNode{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

// Search searches for a value in the BST and returns true if found
func (bst *BST) Search(value int) bool {
	return bst.root.search(value)
}

// search is a helper function to search for a value in the tree
func (n *BSTNode) search(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	if value < n.value {
		return n.left.search(value)
	}
	return n.right.search(value)
}

// InOrder prints the values of the tree in in-order traversal
func (bst *BST) InOrder() {
	if bst.root != nil {
		bst.root.inOrder()
	}
}

// inOrder is a helper function to perform in-order traversal
func (n *BSTNode) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Print(n.value)
	n.right.inOrder()
}
