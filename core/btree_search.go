package core

func BTreeSearch(
	pager *Pager,
	rootID uint64,
	key []byte,
) (
	*Page,
	bool,
) {

	page,
		ok :=
		pager.Get(
			rootID,
		)

	if !ok {
		return nil,
			false
	}

	if PageType(page) != PageTypeInternal {
		return page,
			true
	}

	childID,
		ok :=
		InternalSearch(
			page,
			key,
		)

	if !ok {
		return nil,
			false
	}

	return pager.Get(
		childID,
	)
}
