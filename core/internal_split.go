package core

func InternalSplit(
	left *Page,
	right *Page,
) (
	InternalSplitResult,
	bool,
) {

	count := PageSlotCount(left)

	if count < 2 {
		return InternalSplitResult{}, false
	}

	midpoint := count / 2

	var entries [512]InternalEntry

	for i := uint16(0); i < count; i++ {

		e, ok :=
			ReadInternalEntry(
				left,
				i,
			)

		if !ok {
			return InternalSplitResult{}, false
		}

		entries[i] = e
	}

	separator := entries[midpoint]

	leftID := PageID(left)
	rightID := PageID(right)

	oldLeftChild :=
		LeftChild(left)

	InitInternalPage(
		left,
		leftID,
	)

	SetLeftChild(
		left,
		oldLeftChild,
	)

	InitInternalPage(
		right,
		rightID,
	)

	SetLeftChild(
		right,
		separator.ChildPage,
	)

	for i := uint16(0); i < midpoint; i++ {

		if !InternalInsert(
			left,
			entries[i].Key,
			entries[i].ChildPage,
		) {
			return InternalSplitResult{}, false
		}
	}

	for i := midpoint + 1; i < count; i++ {

		if !InternalInsert(
			right,
			entries[i].Key,
			entries[i].ChildPage,
		) {
			return InternalSplitResult{}, false
		}
	}

	return InternalSplitResult{
		SeparatorKey: separator.Key,
	}, true
}
