package core

func BTreeInsert(
	pager *Pager,
	root *Page,
	key []byte,
	value []byte,
) bool {

	result,
		ok :=
		InsertRecursive(
			pager,
			PageID(root),
			key,
			value,
		)

	if !ok {
		return false
	}

	if result.Split {

		if !HandleRootSplit(
			root,
			result,
		) {
			return false
		}
	}

	return true
}
