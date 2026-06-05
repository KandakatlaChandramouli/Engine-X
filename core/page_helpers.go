package core

import "encoding/binary"

func PageSlotCount(
	p *Page,
) uint16 {

	return binary.LittleEndian.Uint16(
		p.Data[offSlotCount : offSlotCount+2],
	)
}
