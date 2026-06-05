package pagetable

import "sync"

type Entry struct {
	PageID uint64
	LSN    uint64
}

type Table struct {
	mu      sync.RWMutex
	entries map[uint64]Entry
}

func New(capacity int) *Table {

	if capacity < 1 {
		capacity = 1
	}

	return &Table{
		entries: make(
			map[uint64]Entry,
			capacity,
		),
	}
}

func (t *Table) Put(
	pageID uint64,
	lsn uint64,
) {

	t.mu.Lock()

	t.entries[pageID] = Entry{
		PageID: pageID,
		LSN:    lsn,
	}

	t.mu.Unlock()
}

func (t *Table) Get(
	pageID uint64,
) (
	Entry,
	bool,
) {

	t.mu.RLock()

	v,
		ok :=
		t.entries[pageID]

	t.mu.RUnlock()

	return v, ok
}

func (t *Table) Count() int {

	t.mu.RLock()

	n := len(t.entries)

	t.mu.RUnlock()

	return n
}
