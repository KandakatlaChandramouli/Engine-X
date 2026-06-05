package core

func BTreeGet(
	pager *Pager,
	rootID uint64,
	key []byte,
) (
	[]byte,
	bool,
) {

	leaf,
		ok :=
		BTreeSearch(
			pager,
			rootID,
			key,
		)

	if !ok {
		return nil,
			false
	}

	pos :=
		LeafSearch(
			leaf,
			key,
		)

	if pos < 0 {
		return nil,
			false
	}

	_,
		value,
		ok :=
		Get(
			leaf,
			uint16(pos),
		)

	if !ok {
		return nil,
			false
	}

	return value,
		true
}
