package core

func InternalEntryCount(
	p *Page,
) uint16 {

	count := uint16(0)

	for {

		_,
			ok :=
			ReadInternalEntry(
				p,
				count,
			)

		if !ok {
			break
		}

		count++
	}

	return count
}
