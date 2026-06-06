package storage

import "testing"

func TestBackgroundFlusher(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Add(2, 200)

	f := NewBackgroundFlusher(dpt)

	pages := f.Flush()

	if len(pages) != 2 {
		t.Fatal()
	}
}

func TestBackgroundFlusherEmpty(t *testing.T) {
	dpt := NewDirtyPageTable()

	f := NewBackgroundFlusher(dpt)

	pages := f.Flush()

	if len(pages) != 0 {
		t.Fatal()
	}
}

func TestBackgroundFlusherAfterRemove(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Remove(1)

	f := NewBackgroundFlusher(dpt)

	pages := f.Flush()

	if len(pages) != 0 {
		t.Fatal()
	}
}
