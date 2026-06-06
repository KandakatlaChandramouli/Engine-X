package storage

type BufferManager interface {
	FetchPage(id PageID) (*Page, error)
	UnpinPage(id PageID) error
	NewPage() (*Page, error)
}
