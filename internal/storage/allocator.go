package storage

type Allocator struct {
	next PageID
}

func NewAllocator() *Allocator {
	return &Allocator{
		next: 1,
	}
}

func (a *Allocator) Allocate() PageID {
	id := a.next
	a.next++
	return id
}
