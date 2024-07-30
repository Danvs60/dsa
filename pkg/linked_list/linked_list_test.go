package main

import (
	"testing"
)

// TestNewLinkedList tests the NewLinkedList function.
func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll.head != nil {
		t.Errorf("NewLinkedList: expected head to be nil, got %v", ll.head)
	}
}

// TestInsertHead tests the InsertHead function using table-driven tests.
func TestInsertHead(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{"Insert single value", []int{10}, []int{10}},
		{"Insert multiple values", []int{20, 30, 40}, []int{40, 30, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()
			for _, val := range tt.values {
				ll.InsertHead(val)
			}
			if got := ll.GetValues(); !equal(t, got, tt.expected) {
				t.Errorf("InsertHead: expected %v, got %v", tt.expected, got)
			}
		})
	}
}

// TestInsertTail tests the InsertTail function using table-driven tests.
func TestInsertTail(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{"Insert single value", []int{10}, []int{10}},
		{"Insert multiple values", []int{20, 30, 40}, []int{20, 30, 40}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()
			for _, val := range tt.values {
				ll.InsertTail(val)
			}
			if got := ll.GetValues(); !equal(t, got, tt.expected) {
				t.Errorf("InsertTail: expected %v, got %v", tt.expected, got)
			}
		})
	}
}

// TestGet tests the Get function using table-driven tests.
func TestGet(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertHead(10)
	ll.InsertHead(20)
	ll.InsertTail(30)

	tests := []struct {
		name     string
		index    int
		expected int
	}{
		{"Get first element", 0, 20},
		{"Get second element", 1, 10},
		{"Get third element", 2, 30},
		{"Get out of bounds", 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ll.Get(tt.index); got != tt.expected {
				t.Errorf("Get(%d): expected %d, got %d", tt.index, tt.expected, got)
			}
		})
	}
}

// TestRemove tests the Remove function using table-driven tests.
func TestRemove(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertTail(10)
	ll.InsertTail(20)
	ll.InsertTail(30)

	tests := []struct {
		name     string
		index    int
		expected []int
		result   bool
	}{
		{"Remove middle element", 1, []int{10, 30}, true},
		{"Remove out of bounds", 10, []int{10, 30}, false},
		{"Remove first element", 0, []int{30}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ll.Remove(tt.index)
			if res != tt.result {
				t.Errorf("Remove(%d): expected result %v, got %v", tt.index, tt.result, res)
			}
			if got := ll.GetValues(); !equal(t, got, tt.expected) {
				t.Errorf("Remove(%d): expected %v, got %v", tt.index, tt.expected, got)
			}
		})
	}
}

// TestGetValues tests the GetValues function using table-driven tests.
func TestGetValues(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{"Get values from empty list", []int{}, []int{}},
		{"Get values from non-empty list", []int{10, 20, 30}, []int{10, 20, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()
			for _, val := range tt.values {
				ll.InsertTail(val)
			}
			if got := ll.GetValues(); !equal(t, got, tt.expected) {
				t.Errorf("GetValues: expected %v, got %v", tt.expected, got)
			}
		})
	}
}

// Helper function to compare slices for equality
func equal(t *testing.T, a, b []int) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
