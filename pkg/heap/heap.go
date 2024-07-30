package heap

import (
	"errors"
	"fmt"
)

var (
	ErrHeapEmpty                  = errors.New("heap is empty")
	ErrIndexOutOfBounds           = errors.New("index out of bounds")
	ErrNewValueGreaterThanCurrent = errors.New("new value is greater than current")
)

// Heap structure
type Heap struct {
	arr      []int
	size     int
	capacity int
}

// NewHeap creates a new heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		arr:      make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// Insert adds an element to the heap
func (h *Heap) Insert(key int) {
	if h.size == h.capacity {
		fmt.Print("Heap full")
		return
	}

	i := h.size
	h.arr[i] = key // add to end of list
	h.size++       // make space for the next key

	// heapify-up (bubble-up) (MinHeap) (sift-up)
	// used after insertion
	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		// swap parent and child
		h.arr[i], h.arr[parent(i)] = h.arr[parent(i)], h.arr[i]
		i = parent(i)
	}
}

// ExtractMin removes and returns the minimum element from the heap
func (h *Heap) ExtractMin() (int, error) {
	if h.size == 0 {
		return -1, ErrHeapEmpty
	}

	minKey := h.arr[0]
	h.arr[0] = h.arr[h.size-1] // move last key to root to enable bubble-up

	// resize and clean up duplicate
	h.size--
	h.arr = h.arr[:h.size]
	h.MinHeapify(0)

	return minKey, nil
}

// MinHeapify corrects a single violation of the heap property in a subtree's root.
// Also known as bubble-down, sift-down.
// Used when extracting the root (minimum element in a min-heap) or replacing the root
func (h *Heap) MinHeapify(i int) {
	smallest := i

	// left child comparison
	if left(i) < h.size && h.arr[left(i)] < h.arr[smallest] {
		smallest = left(i)
	}

	// right child comparison
	if right(i) < h.size && h.arr[right(i)] < h.arr[smallest] {
		smallest = right(i)
	}

	// recurse swap smallest until heap condition is fulfilled
	if smallest != i {
		h.arr[smallest], h.arr[i] = h.arr[i], h.arr[smallest]
		h.MinHeapify(smallest)
	}
}

// GetMin returns the minimum element from the heap without removing it (peek)
func (h *Heap) GetMin() (int, error) {
	if h.size == 0 {
		return -1, ErrHeapEmpty
	}
	return h.arr[0], nil
}

// DecreaseKey decreases value of key at index i to newVal
func (h *Heap) DecreaseKey(i int, newVal int) error {
	if i < 0 || i >= h.size {
		return ErrIndexOutOfBounds
	}

	if newVal > h.arr[i] {
		return ErrNewValueGreaterThanCurrent
	}

	h.arr[i] = newVal

	// heapify-up - we decreased a key's value somewhere in the tree
	// so bubble it up to keep heap property
	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		// swap with parent if parent is larger
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i) // swap until parent is smaller (thus, heap property is preserved)
	}

	return nil
}

// DeleteKey removes a key at index i
func (h *Heap) DeleteKey(i int) error {
	if h.size == 0 {
		return ErrHeapEmpty
	}
	if i < 0 || i >= h.size {
		return ErrIndexOutOfBounds
	}

	// when key to delete is not the last one
	if i < h.size-1 {
		h.arr[i] = h.arr[h.size-1]
		h.arr = h.arr[:h.size-1]

		h.MinHeapify(i)
	} else {
		// deleting last element
		h.arr = h.arr[:h.size]
	}
	h.size--

	return nil
}

// parent helper function returns the parent index of a heap
func parent(i int) int {
	return (i - 1) / 2
}

// left helper function returns the left child index of a heap
func left(i int) int {
	return 2*i + 1
}

// right helper function returns the right child index of a heap
func right(i int) int {
	return 2*i + 2
}
