package core

import "encoding/binary"

func ReadInternalEntry(
	p *Page,
	slot uint16,
) (
	InternalEntry,
	bool,
) {

	key,
		val,
		ok :=
		Get(
			p,
			slot,
		)

	if !ok {
		return InternalEntry{},
			false
	}

	if len(val) != 8 {
		return InternalEntry{},
			false
	}

	child :=
		binary.LittleEndian.Uint64(
			val,
		)

	return InternalEntry{
		Key:       key,
		ChildPage: child,
	}, true
}
