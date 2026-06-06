package storage

type Flusher struct {
	wal WAL
}

func NewFlusher(w WAL) *Flusher {
	return &Flusher{
		wal: w,
	}
}

func (f *Flusher) FlushPage(pageLSN uint64) error {
	return f.wal.Flush(pageLSN)
}
