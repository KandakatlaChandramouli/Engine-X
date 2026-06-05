package core

func LeafSplit(
	left *Page,
	right *Page,
) (
	SplitResult,
	bool,
) {

	count :=
		LeafEntryCount(
			left,
		)

	if count < 2 {
		return SplitResult{},
			false
	}

	midpoint :=
		count / 2

	var entries [512]Entry

	for i := uint16(0); i < count; i++ {

		e,
			ok :=
			ReadEntry(
				left,
				i,
			)

		if !ok {
			return SplitResult{},
				false
		}

		entries[i] = e
	}

	leftID :=
		PageID(
			left,
		)

	rightID :=
		PageID(
			right,
		)

	InitPage(
		left,
		leftID,
	)

	InitPage(
		right,
		rightID,
	)

	SetPageType(
		left,
		PageTypeLeaf,
	)

	SetPageType(
		right,
		PageTypeLeaf,
	)

	for i := uint16(0); i < midpoint; i++ {

		if !LeafInsert(
			left,
			entries[i].Key,
			entries[i].Val,
		) {
			return SplitResult{},
				false
		}
	}

	for i := midpoint; i < count; i++ {

		if !LeafInsert(
			right,
			entries[i].Key,
			entries[i].Val,
		) {
			return SplitResult{},
				false
		}
	}

	separator,
		_,
		ok :=
		Get(
			right,
			0,
		)

	if !ok {
		return SplitResult{},
			false
	}

	return SplitResult{
		SeparatorKey: separator,
	}, true
}
