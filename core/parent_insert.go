package core

func InsertIntoParent(
	parent *Page,
	separator []byte,
	rightChild uint64,
) bool {

	return InternalInsert(
		parent,
		separator,
		rightChild,
	)
}
