package storage

type DirtyPage struct {
	PageID uint64
	Dirty  bool
	LSN    uint64
}

func (p *DirtyPage) MarkDirty(lsn uint64) {
	p.Dirty = true
	p.LSN = lsn
}

func (p *DirtyPage) MarkClean() {
	p.Dirty = false
}

func (p *DirtyPage) IsDirty() bool {
	return p.Dirty
}
