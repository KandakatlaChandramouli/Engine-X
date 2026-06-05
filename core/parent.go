package core

import "encoding/binary"

const (
	offParentPageID = HeaderSize + 8
)

func ParentPageID(
	p *Page,
) uint64 {

	return binary.LittleEndian.Uint64(
		p.Data[offParentPageID : offParentPageID+8],
	)
}

func SetParentPageID(
	p *Page,
	id uint64,
) {

	binary.LittleEndian.PutUint64(
		p.Data[offParentPageID:offParentPageID+8],
		id,
	)
}
