package core

import "encoding/binary"

const (
	offLeftChild = HeaderSize
)

func LeftChild(
	p *Page,
) uint64 {

	return binary.LittleEndian.Uint64(
		p.Data[offLeftChild : offLeftChild+8],
	)
}

func SetLeftChild(
	p *Page,
	child uint64,
) {

	binary.LittleEndian.PutUint64(
		p.Data[offLeftChild:offLeftChild+8],
		child,
	)
}
