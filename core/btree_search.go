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

	return page,
		true
}
