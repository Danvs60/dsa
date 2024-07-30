package main

import "errors"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

// NewLinkedList intializes an empty linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Get returns the value of the i-th node
func (list *LinkedList) Get(i int) (int, error) {
	current := list.head
	for range i {
		if current == nil {
			return -1, errors.New("index out of bounds")
		}
		current = current.next
	}
	if current == nil {
		return -1, errors.New("index out of bounds")
	}
	return current.value, nil
}

// InsertHead inserts a node with value `val` at the list's head
func (list *LinkedList) InsertHead(val int) {
	newHead := &Node{value: val, next: list.head}
	list.head = newHead
}

// InsertTail inserts a node with value `val` at the list's end
func (list *LinkedList) InsertTail(val int) {
	newTail := &Node{value: val}
	if list.head == nil {
		list.head = newTail
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newTail
	}
}

// Remove removes the i-th node. If out out bounds return false, true otherwise
func (list *LinkedList) Remove(i int) bool {
	current := list.head
	if current == nil {
		return false
	}
	if i == 0 {
		list.head = current.next
		return true
	}

	for range i-1 {
		if current == nil {
			return false
		}
		current = current.next
	}
	if current.next == nil {
		return false
	}
	current.next = current.next.next
	return true
}

// GetValues returns a slice of values in the list
func (list *LinkedList) GetValues() []int {
	var values []int
	current := list.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}
