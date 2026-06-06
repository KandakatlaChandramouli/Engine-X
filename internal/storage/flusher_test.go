package storage

import "testing"

type MockWAL struct {
	Flushed uint64
}

func (m *MockWAL) Flush(lsn uint64) error {
	m.Flushed = lsn
	return nil
}

func TestFlushBeforePage(t *testing.T) {
	w := &MockWAL{}

	f := NewFlusher(w)

	err := f.FlushPage(100)

	if err != nil {
		t.Fatal()
	}

	if w.Flushed != 100 {
		t.Fatal()
	}
}
