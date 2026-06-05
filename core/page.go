package core

import "encoding/binary"

const (
	PageSize   = 4096
	HeaderSize = 32
	SlotSize   = 8
)

const (
	offPageID    = 0
	offLSN       = 8
	offFreeStart = 16
	offFreeEnd   = 18
	offSlotCount = 20
	offFlags     = 22
)

type Page struct {
	Data [PageSize]byte
}

func InitPage(p *Page, id uint64) {
	binary.LittleEndian.PutUint64(
		p.Data[offPageID:offPageID+8],
		id,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offLSN:offLSN+8],
		0,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFreeStart:offFreeStart+2],
		HeaderSize,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFreeEnd:offFreeEnd+2],
		PageSize,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offSlotCount:offSlotCount+2],
		0,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFlags:offFlags+2],
		0,
	)
}

func PageID(p *Page) uint64 {
	return binary.LittleEndian.Uint64(
		p.Data[offPageID : offPageID+8],
	)
}

func SetLSN(p *Page, lsn uint64) {
	binary.LittleEndian.PutUint64(
		p.Data[offLSN:offLSN+8],
		lsn,
	)
}

func Insert(
	p *Page,
	key []byte,
	val []byte,
) bool {

	freeStart :=
		binary.LittleEndian.Uint16(
			p.Data[offFreeStart : offFreeStart+2],
		)

	freeEnd :=
		binary.LittleEndian.Uint16(
			p.Data[offFreeEnd : offFreeEnd+2],
		)

	slotCount :=
		binary.LittleEndian.Uint16(
			p.Data[offSlotCount : offSlotCount+2],
		)

	payloadSize :=
		len(key) + len(val)

	required :=
		payloadSize + SlotSize

	if int(freeEnd-freeStart) < required {
		return false
	}

	payloadOffset :=
		freeEnd - uint16(payloadSize)

	copy(
		p.Data[payloadOffset:payloadOffset+uint16(len(key))],
		key,
	)

	copy(
		p.Data[payloadOffset+uint16(len(key)):payloadOffset+uint16(payloadSize)],
		val,
	)

	slotPos :=
		HeaderSize + int(slotCount)*SlotSize

	binary.LittleEndian.PutUint16(
		p.Data[slotPos:slotPos+2],
		payloadOffset,
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+2:slotPos+4],
		uint16(len(key)),
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+4:slotPos+6],
		uint16(len(val)),
	)

	binary.LittleEndian.PutUint16(
		p.Data[slotPos+6:slotPos+8],
		0,
	)

	slotCount++

	binary.LittleEndian.PutUint16(
		p.Data[offSlotCount:offSlotCount+2],
		slotCount,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFreeStart:offFreeStart+2],
		freeStart+SlotSize,
	)

	binary.LittleEndian.PutUint16(
		p.Data[offFreeEnd:offFreeEnd+2],
		payloadOffset,
	)

	return true
}

const (
	PageTypeMeta uint16 = 1

	PageTypeFreelist uint16 = 2

	PageTypeLeaf uint16 = 3

	PageTypeInternal uint16 = 4

	PageTypeOverflow uint16 = 5
)

func SetPageType(
	p *Page,
	typ uint16,
) {

	binary.LittleEndian.PutUint16(
		p.Data[offFlags:offFlags+2],
		typ,
	)
}

func PageType(
	p *Page,
) uint16 {

	return binary.LittleEndian.Uint16(
		p.Data[offFlags : offFlags+2],
	)
}
