package core

import "encoding/binary"

func InitInternalPage(
	p *Page,
	id uint64,
) {

	InitPage(
		p,
		id,
	)

	SetPageType(
		p,
		PageTypeInternal,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFreeStart:offFreeStart+2],
		InternalHeaderSize,
	)
}
