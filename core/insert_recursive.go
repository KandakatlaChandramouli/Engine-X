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

	ok =
		LeafInsert(
			leaf,
			key,
			value,
		)

	if ok {
		return InsertRecursiveResult{},
			true
	}

	return InsertRecursiveResult{
		Split: true,
	}, true
}
