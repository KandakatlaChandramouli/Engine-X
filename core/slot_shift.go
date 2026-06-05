package core

func ShiftSlotsRight(
	p *Page,
	start uint16,
) {

	slotCount :=
		PageSlotCount(
			p,
		)

	if slotCount == 0 {
		return
	}

	for i := slotCount; i > start; i-- {

		CopySlot(
			p,
			i,
			i-1,
		)
	}
}
