package core

import "encoding/binary"

type Slot struct {
	Offset uint16
	KeyLen uint16
	ValLen uint16
	Flags  uint16
}

func ReadSlot(
	p *Page,
	index uint16,
) Slot {

	slotPos :=
		SlotBase(p) +
			int(index)*SlotSize

	return Slot{
		Offset: binary.LittleEndian.Uint16(
			p.Data[slotPos : slotPos+2],
		),
		KeyLen: binary.LittleEndian.Uint16(
			p.Data[slotPos+2 : slotPos+4],
		),
		ValLen: binary.LittleEndian.Uint16(
			p.Data[slotPos+4 : slotPos+6],
		),
		Flags: binary.LittleEndian.Uint16(
			p.Data[slotPos+6 : slotPos+8],
		),
	}
}

func WriteSlot(
	p *Page,
	index uint16,
	s Slot,
) {

	slotPos :=
		SlotBase(p) +
			int(index)*SlotSize

	binary.LittleEndian.PutUint16(
		p.Data[slotPos:slotPos+2],
		s.Offset,
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+2:slotPos+4],
		s.KeyLen,
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+4:slotPos+6],
		s.ValLen,
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+6:slotPos+8],
		s.Flags,
	)
}
