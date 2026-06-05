package core

func InternalSearch(
	p *Page,
	key []byte,
) (
	uint64,
	bool,
) {

	idx :=
		LeafSearch(
			p,
			key,
		)

	if idx < 0 {

		insertPos :=
			-(idx + 1)

		if insertPos == 0 {

			return LeftChild(
				p,
			), true
		}

		idx = insertPos - 1
	}

	e,
		ok :=
		ReadInternalEntry(
			p,
			uint16(idx),
		)

	if !ok {
		return 0,
			false
	}

	return e.ChildPage,
		true
}
