package core

import (
	"encoding/binary"
	"hash/crc32"
)

const (
	MetaMagic   uint64 = 0x454E47494E4558
	MetaVersion uint64 = 1

	MetaSize = 64

	offMetaMagic      = 0
	offMetaVersion    = 8
	offMetaRootPageID = 16
	offMetaFreelistID = 24
	offMetaLastPageID = 32
	offMetaTXID       = 40
	offMetaCRC32      = 48
)

type Meta struct {
	Magic      uint64
	Version    uint64
	RootPageID uint64
	FreelistID uint64
	LastPageID uint64
	TXID       uint64
	CRC32      uint32
}

func InitMeta(
	p *Page,
) {

	binary.LittleEndian.PutUint64(
		p.Data[offMetaMagic:offMetaMagic+8],
		MetaMagic,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offMetaVersion:offMetaVersion+8],
		MetaVersion,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offMetaRootPageID:offMetaRootPageID+8],
		3,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offMetaFreelistID:offMetaFreelistID+8],
		2,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offMetaLastPageID:offMetaLastPageID+8],
		3,
	)

	binary.LittleEndian.PutUint64(
		p.Data[offMetaTXID:offMetaTXID+8],
		0,
	)

	UpdateMetaCRC(p)
}

func UpdateMetaCRC(
	p *Page,
) {

	crc :=
		crc32.ChecksumIEEE(
			p.Data[:48],
		)

	binary.LittleEndian.PutUint32(
		p.Data[offMetaCRC32:offMetaCRC32+4],
		crc,
	)
}

func ValidateMeta(
	p *Page,
) bool {

	if binary.LittleEndian.Uint64(
		p.Data[offMetaMagic:offMetaMagic+8],
	) != MetaMagic {
		return false
	}

	crc :=
		crc32.ChecksumIEEE(
			p.Data[:48],
		)

	stored :=
		binary.LittleEndian.Uint32(
			p.Data[offMetaCRC32 : offMetaCRC32+4],
		)

	return crc == stored
}

func RootPageID(
	p *Page,
) uint64 {

	return binary.LittleEndian.Uint64(
		p.Data[offMetaRootPageID : offMetaRootPageID+8],
	)
}

func LastPageID(
	p *Page,
) uint64 {

	return binary.LittleEndian.Uint64(
		p.Data[offMetaLastPageID : offMetaLastPageID+8],
	)
}

func SetLastPageID(
	p *Page,
	id uint64,
) {

	binary.LittleEndian.PutUint64(
		p.Data[offMetaLastPageID:offMetaLastPageID+8],
		id,
	)

	UpdateMetaCRC(p)
}
