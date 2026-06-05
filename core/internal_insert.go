package core

import "encoding/binary"

func InternalInsert(
	p *Page,
	key []byte,
	child uint64,
) bool {

	var buf [8]byte

	binary.LittleEndian.PutUint64(
		buf[:],
		child,
	)

	return LeafInsert(
		p,
		key,
		buf[:],
	)
}
