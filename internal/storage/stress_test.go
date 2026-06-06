package storage

import "testing"

func TestDirtyPageStress(t *testing.T) {
	dpt := NewDirtyPageTable()

	for i := uint64(0); i < 1000; i++ {
		dpt.Add(i, i*10)
	}

	if dpt.Len() != 1000 {
		t.Fatal()
	}
}

func TestFlushStress(t *testing.T) {
	dpt := NewDirtyPageTable()

	for i := uint64(0); i < 1000; i++ {
		dpt.Add(i, i)
	}

	pages := FlushCandidates(dpt)

	if len(pages) != 1000 {
		t.Fatal()
	}
}

func TestRecoveryStress(t *testing.T) {
	dpt := NewDirtyPageTable()

	for i := uint64(1); i <= 1000; i++ {
		dpt.Add(i, i)
	}

	r := NewRecoveryHook(dpt)

	if r.RecoveryStartLSN() != 1 {
		t.Fatal()
	}
}
