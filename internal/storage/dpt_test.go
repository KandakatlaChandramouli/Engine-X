package storage

import "testing"

func TestDPTAdd(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)

	if dpt.Len() != 1 {
		t.Fatal()
	}
}

func TestDPTRecLSN(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)

	lsn, ok := dpt.RecLSN(1)

	if !ok {
		t.Fatal()
	}

	if lsn != 100 {
		t.Fatal()
	}
}

func TestDPTRemove(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Remove(1)

	if dpt.Len() != 0 {
		t.Fatal()
	}
}

func TestDPTKeepsFirstRecLSN(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 100)
	dpt.Add(1, 200)

	lsn, _ := dpt.RecLSN(1)

	if lsn != 100 {
		t.Fatal()
	}
}
