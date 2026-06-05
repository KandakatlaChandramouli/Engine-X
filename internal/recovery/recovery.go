package recovery

import (
	"encoding/binary"
	"errors"
	"io"
	"os"

	"engine-x/wal"
)

type Mutation struct {
	PageID uint64
	Key    []byte
	Value  []byte
	LSN    uint64
}

func Replay(
	path string,
	apply func(Mutation) error,
) (uint64, error) {

	f, err := os.Open(path)

	if err != nil {

		if errors.Is(
			err,
			os.ErrNotExist,
		) {
			return 0, nil
		}

		return 0, err
	}

	defer f.Close()

	var recovered uint64

	for {

		hdr,
			payload,
			err :=
			wal.ReadRecord(f)

		if err == io.EOF {
			break
		}

		if err == io.ErrUnexpectedEOF {
			break
		}

		if err != nil {
			break
		}

		if len(payload) < 12 {
			break
		}

		pageID :=
			binary.LittleEndian.Uint64(
				payload[0:8],
			)

		keyLen :=
			binary.LittleEndian.Uint16(
				payload[8:10],
			)

		valLen :=
			binary.LittleEndian.Uint16(
				payload[10:12],
			)

		expected :=
			12 +
				int(keyLen) +
				int(valLen)

		if expected != len(payload) {
			break
		}

		keyStart := 12
		keyEnd := keyStart + int(keyLen)

		valStart := keyEnd
		valEnd := valStart + int(valLen)

		m := Mutation{
			PageID: pageID,
			Key:    payload[keyStart:keyEnd],
			Value:  payload[valStart:valEnd],
			LSN:    hdr.LSN,
		}

		if err := apply(m); err != nil {
			return recovered, err
		}

		recovered++
	}

	return recovered, nil
}
