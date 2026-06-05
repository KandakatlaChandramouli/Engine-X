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

func HandleParentOverflow(
	parent *Page,
	sibling *Page,
) (
	InternalSplitResult,
	bool,
) {

	return InternalSplit(
		parent,
		sibling,
	)
}
