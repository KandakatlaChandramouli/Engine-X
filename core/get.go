package core

import "encoding/binary"

func Get(
	p *Page,
	slot uint16,
) (
	[]byte,
	[]byte,
	bool,
) {

	slotCount :=
		binary.LittleEndian.Uint16(
			p.Data[offSlotCount : offSlotCount+2],
		)

	if slot >= slotCount {
		return nil, nil, false
	}

	slotPos :=
		SlotBase(p) +
			int(slot)*SlotSize

	payloadOffset :=
		binary.LittleEndian.Uint16(
			p.Data[slotPos : slotPos+2],
		)

	keyLen :=
		binary.LittleEndian.Uint16(
			p.Data[slotPos+2 : slotPos+4],
		)

	valLen :=
		binary.LittleEndian.Uint16(
			p.Data[slotPos+4 : slotPos+6],
		)

	payloadStart :=
		int(payloadOffset)

	keyStart :=
		payloadStart

	keyEnd :=
		keyStart + int(keyLen)

	valStart :=
		keyEnd

	valEnd :=
		valStart + int(valLen)

	if payloadStart < HeaderSize {
		return nil, nil, false
	}

	if keyEnd > PageSize {
		return nil, nil, false
	}

	if valEnd > PageSize {
		return nil, nil, false
	}

	key :=
		p.Data[keyStart:keyEnd]

	val :=
		p.Data[valStart:valEnd]

	return key,
		val,
		true
}
