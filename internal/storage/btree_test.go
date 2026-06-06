package storage

import "testing"

type MockBuffer struct{}

func (m *MockBuffer) FetchPage(id PageID) (*Page, error) {
	return &Page{ID: id}, nil
}

func (m *MockBuffer) UnpinPage(id PageID) error {
	return nil
}

func (m *MockBuffer) NewPage() (*Page, error) {
	return &Page{}, nil
}

func TestBTreeCreation(t *testing.T) {
	buf := &MockBuffer{}

	tree := NewBTree(buf)

	if tree == nil {
		t.Fatal()
	}
}
