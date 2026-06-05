package core

func CopySlot(
	p *Page,
	dst uint16,
	src uint16,
) {

	WriteSlot(
		p,
		dst,
		ReadSlot(
			p,
			src,
		),
	)
}
