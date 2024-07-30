package heap

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	capacity := 2
	got := NewHeap(capacity)
	want := &Heap{
		arr:      make([]int, 2),
		size:     0,
		capacity: capacity,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestInsert(t *testing.T) {
	h := NewHeap(3)
	h.Insert(2)
	h.Insert(3)
	h.Insert(1)

	want := []int{1, 3, 2}
	got := h.arr

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestExtractMin(t *testing.T) {
	h := NewHeap(5)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)

	wantMin := 1
	gotMin, err := h.ExtractMin()
	if err != nil {
		t.Errorf("ExtractMin() failed with error %v", err)
	}

	if gotMin != wantMin {
		t.Errorf("want min: %v, got %v", wantMin, gotMin)
	}

	gotArr := h.arr
	wantArr := []int{2, 4, 3, 5}
	if !reflect.DeepEqual(gotArr, wantArr) {
		t.Errorf("got %#v, want %#v", gotArr, wantArr)
	}
}

func TestMinHeapify(t *testing.T) {
	h := NewHeap(4)
	h.arr[0] = 5
	h.arr[1] = 2
	h.arr[2] = 3
	h.arr[3] = 4
	h.size = 4

	h.MinHeapify(0)

	got := h.arr[:h.size] // Slice to the current size of the heap
	want := []int{2, 4, 3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestGetMin(t *testing.T) {
	tests := []struct {
		name        string
		heapArr     []int
		expectedMin int
		wantErr     bool
	}{
		{
			name:        "Empty heap",
			heapArr:     make([]int, 0),
			expectedMin: -1,
			wantErr:     true,
		},
		{
			name:        "Valid heap",
			heapArr:     []int{1, 2, 3},
			expectedMin: 1,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(len(tt.heapArr))
			h.arr = tt.heapArr
			h.size = len(h.arr)
			gotMin, err := h.GetMin()

			if tt.wantErr && err == nil {
				t.Error("Expected error, but didn't get one")
			}

			if err != nil && gotMin != tt.expectedMin {
				t.Errorf("got min %v, expected %v", gotMin, tt.expectedMin)
			}

			if !reflect.DeepEqual(h.arr, tt.heapArr) {
				t.Errorf("Heap was changed, even though using peek method")
			}
		})
	}
}

func TestDeleteKey(t *testing.T) {
	h := NewHeap(5)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3) // delete this i:2
	h.Insert(5)
	h.Insert(6)

	_ = h.DeleteKey(2)

	wantArr := []int{1,2,6,5}
	if !reflect.DeepEqual(h.arr, wantArr) {
		t.Errorf("want %v, got %v", wantArr, h.arr)
	}

	err := h.DeleteKey(7)
	if err == nil {
		t.Error("expected error but didn't get one")
	}
}

func TestDecreaseKey(t *testing.T) {
	tests := []struct {
		name          string
		inputArr      []int
		wantArr       []int
		decreaseIndex int
		newVal        int
		wantErr       bool
	}{
		{
			name:          "valid decrease",
			inputArr:      []int{2, 3, 4},
			decreaseIndex: 1,
			newVal:        1,
			wantArr:       []int{1, 2, 4},
			wantErr:       false,
		},
		{
			name:          "invalid decrease",
			inputArr:      []int{2, 3, 4},
			decreaseIndex: 1,
			newVal:        4,
			wantArr:       []int{2, 3, 4},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(len(tt.inputArr))
			h.size = len(tt.inputArr)
			h.arr = tt.inputArr

			err := h.DecreaseKey(tt.decreaseIndex, tt.newVal)
			if tt.wantErr && err == nil {
				t.Error("expected error but didn't get one")
			}

			gotArr := h.arr
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("want %v, got %v", tt.wantArr, gotArr)
			}
		})
	}
}

func TestParent(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  int
	}{
		{
			name:  "parent of odd",
			index: 1,
			want:  0,
		},
		{
			name:  "parent of even",
			index: 2,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parent(tt.index)
			if got != tt.want {
				t.Errorf("given input index: %v, got %v, want %v", tt.index, got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  int
	}{
		{
			name:  "left of odd",
			index: 0,
			want:  1,
		},
		{
			name:  "left of even",
			index: 2,
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := left(tt.index)
			if got != tt.want {
				t.Errorf("given input index: %v, got %v, want %v", tt.index, got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  int
	}{
		{
			name:  "right of odd",
			index: 0,
			want:  2,
		},
		{
			name:  "right of even",
			index: 2,
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := right(tt.index)
			if got != tt.want {
				t.Errorf("given input index: %v, got %v, want %v", tt.index, got, tt.want)
			}
		})
	}
}
