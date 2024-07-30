package main

import (
	"errors"
)

// Deque represents a double-ended queue
type Deque struct {
	elements []int
}

// NewDeque creates a new Deque
func NewDeque() *Deque {
	return &Deque{elements: []int{}}
}

// AddFront adds an element to the front of the deque
func (d *Deque) AddFront(value int) {
	d.elements = append([]int{value}, d.elements...)
}

// AddRear adds an element to the rear of the deque
func (d *Deque) AddRear(value int) {
	d.elements = append(d.elements, value)
}

// RemoveFront removes and returns the element from the front
func (d *Deque) RemoveFront() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("Empty")
	}
	front := d.elements[0]
	d.elements = d.elements[1:]
	return front, nil
}

// RemoveRear removes and returns the element from the rear
func (d *Deque) RemoveRear() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("Empty")
	}
	rear := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return rear, nil
}

// IsEmpty checks if the deque is empty
func (d *Deque) IsEmpty() bool {
	return len(d.elements) == 0
}

// Size returns the number of elements in the deque
func (d *Deque) Size() int {
	return len(d.elements)
}

// PeekFront returns the front element without removing it
func (d *Deque) PeekFront() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("empty")
	}
	return d.elements[0], nil
}

// PeekRear returns the rear element without removing it
func (d *Deque) PeekRear() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("empty")
	}
	return d.elements[len(d.elements)-1], nil
}
