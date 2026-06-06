package storage

import "testing"

func TestRecoveryStartLSN(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 300)
	dpt.Add(2, 100)
	dpt.Add(3, 200)

	r := NewRecoveryHook(dpt)

	if r.RecoveryStartLSN() != 100 {
		t.Fatal()
	}
}

func TestRecoveryStartLSNEmpty(t *testing.T) {
	dpt := NewDirtyPageTable()

	r := NewRecoveryHook(dpt)

	if r.RecoveryStartLSN() != 0 {
		t.Fatal()
	}
}

func TestRecoveryStartLSNSingle(t *testing.T) {
	dpt := NewDirtyPageTable()

	dpt.Add(1, 500)

	r := NewRecoveryHook(dpt)

	if r.RecoveryStartLSN() != 500 {
		t.Fatal()
	}
}
