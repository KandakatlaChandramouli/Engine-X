package core

import "encoding/binary"

const (
	offFreelistCount = HeaderSize
	offFreelistItems = HeaderSize + 8
)

func InitFreelist(
	p *Page,
) {

	InitPage(
		p,
		2,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offFreelistCount:offFreelistCount+8],
		0,
	)
}

func FreelistCount(
	p *Page,
) uint64 {

	return binary.LittleEndian.Uint64(
		p.Data[offFreelistCount : offFreelistCount+8],
	)
}

func FreelistPush(
	p *Page,
	pageID uint64,
) bool {

	count :=
		FreelistCount(p)

	offset :=
		offFreelistItems +
			int(count)*8

	if offset+8 > PageSize {
		return false
	}

	binary.LittleEndian.PutUint64(
		p.Data[offset:offset+8],
		pageID,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offFreelistCount:offFreelistCount+8],
		count+1,
	)

	return true
}

func FreelistPop(
	p *Page,
) (
	uint64,
	bool,
) {

	count :=
		FreelistCount(p)

	if count == 0 {
		return 0, false
	}

	last :=
		count - 1

	offset :=
		offFreelistItems +
			int(last)*8

	pageID :=
		binary.LittleEndian.Uint64(
			p.Data[offset : offset+8],
		)

	binary.LittleEndian.PutUint64(
		p.Data[offFreelistCount:offFreelistCount+8],
		last,
	)

	return pageID, true
}
