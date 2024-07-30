package main

import "testing"

func TestHashFunction(t *testing.T) {
	ht := NewHashTable(20)
	want := 5
	got := ht.hashFunction("A")

	if want != got {
		t.Fatalf("got %v, want %v", got, want)
	}
}
