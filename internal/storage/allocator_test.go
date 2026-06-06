package storage

import "testing"

func TestAllocate(t *testing.T) {
	a := NewAllocator()

	id := a.Allocate()

	if id != 1 {
		t.Fatal()
	}
}

func TestAllocateSequence(t *testing.T) {
	a := NewAllocator()

	if a.Allocate() != 1 {
		t.Fatal()
	}

	if a.Allocate() != 2 {
		t.Fatal()
	}

	if a.Allocate() != 3 {
		t.Fatal()
	}
}
