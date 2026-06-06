package storage

type BTree struct {
	buffer BufferManager
}

func NewBTree(buf BufferManager) *BTree {
	return &BTree{
		buffer: buf,
	}
}
