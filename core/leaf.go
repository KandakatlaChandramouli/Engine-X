package core

import "bytes"

func LeafSearch(
	p *Page,
	key []byte,
) int {

	slotCount := int(
		PageSlotCount(p),
	)

	low := 0
	high := slotCount - 1

	for low <= high {

		mid :=
			low +
				(high-low)/2

		k,
			_,
			ok :=
			Get(
				p,
				uint16(mid),
			)

		if !ok {
			return -1
		}

		cmp :=
			bytes.Compare(
				k,
				key,
			)

		if cmp == 0 {
			return mid
		}

		if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -(low + 1)
}

func LeafGet(
	p *Page,
	key []byte,
) (
	[]byte,
	bool,
) {

	idx :=
		LeafSearch(
			p,
			key,
		)

	if idx < 0 {
		return nil, false
	}

	_,
		val,
		ok :=
		Get(
			p,
			uint16(idx),
		)

	if !ok {
		return nil, false
	}

	return val, true
}
