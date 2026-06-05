package core

func SlotBase(
	p *Page,
) int {

	if PageType(p) == PageTypeInternal {
		return InternalHeaderSize
	}

	return HeaderSize
}
