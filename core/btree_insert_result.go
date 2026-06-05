package core

type BTreeInsertResult struct {
	Split        bool
	SeparatorKey []byte
}
