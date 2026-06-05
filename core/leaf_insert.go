package core

func LeafInsert(
	p *Page,
	key []byte,
	val []byte,
) bool {

	pos :=
		LeafSearch(
			p,
			key,
		)

	if pos >= 0 {
		return false
	}

	insertPos :=
		uint16(
			-(pos + 1),
		)

	slotCount :=
		PageSlotCount(
			p,
		)

	if !Insert(
		p,
		key,
		val,
	) {
		return false
	}

	newSlot :=
		ReadSlot(
			p,
			slotCount,
		)

	ShiftSlotsRight(
		p,
		insertPos,
	)

	WriteSlot(
		p,
		insertPos,
		newSlot,
	)

	return true
}
