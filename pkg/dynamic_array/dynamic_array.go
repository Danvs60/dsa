package main

import (
	"errors"
)

type DynamicArray struct {
	data     []int
	size     int
	capacity int
}

const capacityFactor = 2

// NewDynamicArray returns a new dynamic array with the given capacity
func NewDynamicArray(capacity int) *DynamicArray {
	if capacity <= 0 {
		panic("capacity must be greater than 0")
	}
	return &DynamicArray{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (da *DynamicArray) isOutOfBounds(i int) error {
	if i >= da.size {
		return errors.New("index out of bounds")
	}
	return nil
}

// Get returns the element at index i
func (da *DynamicArray) Get(i int) (int, error) {
	if err := da.isOutOfBounds(i); err != nil {
		return 0, err
	}
	return da.data[i], nil
}

// Set assigns the element at index i the value of n
func (da *DynamicArray) Set(i int, n int) error {
	if err := da.isOutOfBounds(i); err != nil {
		return err
	}
	da.data[i] = n
	return nil
}

// PushBack adds an element with value n to the end of the array
func (da *DynamicArray) PushBack(n int) {
	if da.size == da.capacity {
		da.resize()
	}
	da.data[da.size] = n
	da.size++
}

func (da *DynamicArray) resize() {
	da.capacity *= capacityFactor
	newData := make([]int, da.capacity)
	copy(newData, da.data)
	da.data = newData
}

// PopBack removes and returns the element at the end of the array
func (da *DynamicArray) PopBack() (int, error) {
	if da.size == 0 {
		return 0, errors.New("cannot pop, array is empty")
	}
	v := da.data[da.size-1]
	da.size--
	return v, nil
}

// GetSize returns the number of elements in the array
func (da *DynamicArray) GetSize() int {
	return da.size
}

// GetCapacity returns the capacity of the array
func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
