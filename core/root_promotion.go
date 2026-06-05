package core

func PromoteRoot(
	root *Page,
	leftPageID uint64,
	rightPageID uint64,
	separator []byte,
) bool {

	InitInternalPage(
		root,
		PageID(root),
	)

	SetLeftChild(
		root,
		leftPageID,
	)

	return InternalInsert(
		root,
		separator,
		rightPageID,
	)
}
