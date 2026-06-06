package storage

type BackgroundFlusher struct {
	dpt *DirtyPageTable
}

func NewBackgroundFlusher(dpt *DirtyPageTable) *BackgroundFlusher {
	return &BackgroundFlusher{
		dpt: dpt,
	}
}

func (f *BackgroundFlusher) Flush() []uint64 {
	return FlushCandidates(f.dpt)
}
