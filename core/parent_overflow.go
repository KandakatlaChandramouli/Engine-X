package core

func ParentOverflow(
	parent *Page,
	separator []byte,
	rightChild uint64,
) bool {

	return !InternalInsert(
		parent,
		separator,
		rightChild,
	)
}
