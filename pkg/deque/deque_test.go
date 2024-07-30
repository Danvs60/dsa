package main

import (
	"testing"
)

// TestAddFront tests the AddFront method
func TestAddFront(t *testing.T) {
	d := NewDeque()
	d.AddFront(10)
	if d.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", d.Size())
	}
	if value, _ := d.PeekFront(); value != 10 {
		t.Errorf("Expected front value to be 10, got %d", value)
	}
}

// TestAddRear tests the AddRear method
func TestAddRear(t *testing.T) {
	d := NewDeque()
	d.AddRear(20)
	if d.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", d.Size())
	}
	if value, _ := d.PeekRear(); value != 20 {
		t.Errorf("Expected rear value to be 20, got %d", value)
	}
}

// TestRemoveFront tests the RemoveFront method
func TestRemoveFront(t *testing.T) {
	d := NewDeque()
	d.AddFront(30)
	value, err := d.RemoveFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected value to be 30, got %d", value)
	}
	if !d.IsEmpty() {
		t.Errorf("Expected deque to be empty")
	}
}

// TestRemoveRear tests the RemoveRear method
func TestRemoveRear(t *testing.T) {
	d := NewDeque()
	d.AddRear(40)
	value, err := d.RemoveRear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 40 {
		t.Errorf("Expected value to be 40, got %d", value)
	}
	if !d.IsEmpty() {
		t.Errorf("Expected deque to be empty")
	}
}

// TestIsEmpty tests the IsEmpty method
func TestIsEmpty(t *testing.T) {
	d := NewDeque()
	if !d.IsEmpty() {
		t.Errorf("Expected deque to be empty")
	}
	d.AddFront(50)
	if d.IsEmpty() {
		t.Errorf("Expected deque to not be empty")
	}
}

// TestSize tests the Size method
func TestSize(t *testing.T) {
	d := NewDeque()
	if d.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", d.Size())
	}
	d.AddFront(60)
	if d.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", d.Size())
	}
}

// TestPeekFront tests the PeekFront method
func TestPeekFront(t *testing.T) {
	d := NewDeque()
	d.AddFront(70)
	value, err := d.PeekFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 70 {
		t.Errorf("Expected value to be 70, got %d", value)
	}
}

// TestPeekRear tests the PeekRear method
func TestPeekRear(t *testing.T) {
	d := NewDeque()
	d.AddRear(80)
	value, err := d.PeekRear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 80 {
		t.Errorf("Expected value to be 80, got %d", value)
	}
}
