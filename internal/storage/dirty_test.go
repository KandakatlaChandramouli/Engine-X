package storage

import "testing"

func TestDirtyPage(t *testing.T) {
	p := &DirtyPage{}

	p.MarkDirty(100)

	if !p.IsDirty() {
		t.Fatal()
	}

	if p.LSN != 100 {
		t.Fatal()
	}
}

func TestCleanPage(t *testing.T) {
	p := &DirtyPage{}

	p.MarkDirty(100)
	p.MarkClean()

	if p.IsDirty() {
		t.Fatal()
	}
}

func TestLSNUpdate(t *testing.T) {
	p := &DirtyPage{}

	p.MarkDirty(10)
	p.MarkDirty(20)

	if p.LSN != 20 {
		t.Fatal()
	}
}
