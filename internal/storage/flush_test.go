package storage

import "testing"

func TestFlushCandidates(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Add(2, 200)

	pages := FlushCandidates(dpt)

	if len(pages) != 2 {
		t.Fatal()
	}
}

func TestFlushCandidatesEmpty(t *testing.T) {
	dpt := NewDirtyPageTable()

	pages := FlushCandidates(dpt)

	if len(pages) != 0 {
		t.Fatal()
	}
}

func TestFlushCandidatesAfterRemove(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Remove(1)

	pages := FlushCandidates(dpt)

	if len(pages) != 0 {
		t.Fatal()
	}
}
