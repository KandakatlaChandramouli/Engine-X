package core

func InsertRecursive(
	pager *Pager,
	rootID uint64,
	key []byte,
	value []byte,
) (
	InsertRecursiveResult,
	bool,
) {

	path,
		ok :=
		FindInsertionPath(
			pager,
			rootID,
			key,
		)

	if !ok {
		return InsertRecursiveResult{},
			false
	}

	leaf :=
		path.Pages[len(path.Pages)-1]

	if LeafInsert(
		leaf,
		key,
		value,
	) {
		return InsertRecursiveResult{},
			true
	}

	siblingID :=
		pager.AllocatePageID()

	sibling,
		ok :=
		AllocateLeafSibling(
			pager,
			siblingID,
		)

	if !ok {
		return InsertRecursiveResult{},
			false
	}

	split,
		ok :=
		LeafSplit(
			leaf,
			sibling,
		)

	if !ok {
		return InsertRecursiveResult{},
			false
	}

	return InsertRecursiveResult{
		Split:        true,
		SeparatorKey: split.SeparatorKey,
		RightPageID:  PageID(sibling),
	}, true
}
