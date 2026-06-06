package storage

type WAL interface {
	Flush(lsn uint64) error
}
