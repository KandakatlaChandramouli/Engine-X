package storage

type PageID uint64

type Page struct {
	ID   PageID
	Data []byte
}
