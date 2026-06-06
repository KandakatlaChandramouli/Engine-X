package storage

type DirtyPageTable struct {
	pages map[uint64]uint64
}

func NewDirtyPageTable() *DirtyPageTable {
	return &DirtyPageTable{
		pages: make(map[uint64]uint64),
	}
}

func (d *DirtyPageTable) Add(pageID uint64, recLSN uint64) {
	if _, ok := d.pages[pageID]; !ok {
		d.pages[pageID] = recLSN
	}
}

func (d *DirtyPageTable) Remove(pageID uint64) {
	delete(d.pages, pageID)
}

func (d *DirtyPageTable) RecLSN(pageID uint64) (uint64, bool) {
	lsn, ok := d.pages[pageID]
	return lsn, ok
}

func (d *DirtyPageTable) Len() int {
	return len(d.pages)
}
