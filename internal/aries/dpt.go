package aries

type DirtyPageTable struct {
        pages map[uint64]DirtyPageEntry
}

func NewDirtyPageTable() *DirtyPageTable {

        return &DirtyPageTable{
                pages: make(
                        map[uint64]DirtyPageEntry,
                ),
        }
}

func (d *DirtyPageTable) Add(
        pageID uint64,
        recLSN uint64,
) {

        if _, ok := d.pages[pageID]; ok {
                return
        }

        d.pages[pageID] = DirtyPageEntry{
                PageID: pageID,
                RecLSN: recLSN,
        }
}

func (d *DirtyPageTable) Get(
        pageID uint64,
) (
        DirtyPageEntry,
        bool,
) {
        v, ok := d.pages[pageID]
        return v, ok
}

func (d *DirtyPageTable) Count() int {
        return len(d.pages)
}
