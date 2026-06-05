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

	return InsertRecursiveResult{
		Split: true,
		RightPageID: uint64(
			len(pager.pages) + 1,
		),
	}, true
}
