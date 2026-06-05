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
		idx = -(idx + 1)

		if idx >= int(
			PageSlotCount(
				p,
			),
		) {
			idx--
		}
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
