package wal

import (
	"bufio"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
	"sync"
)

const (
	HeaderSize = 24
)

type WAL struct {
	mu   sync.Mutex
	file *os.File
	buf  *bufio.Writer
	lsn  uint64
}

func OpenWAL(
	path string,
) (*WAL, error) {

	f, err :=
		os.OpenFile(
			path,
			os.O_CREATE|
				os.O_RDWR|
				os.O_APPEND,
			0644,
		)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: f,
		buf: bufio.NewWriterSize(
			f,
			1<<20,
		),
	}, nil
}

func (w *WAL) Append(
	txid uint64,
	payload []byte,
) (uint64, error) {

	w.mu.Lock()
	defer w.mu.Unlock()

	w.lsn++

	var hdr [HeaderSize]byte

	binary.LittleEndian.PutUint64(
		hdr[4:12],
		w.lsn,
	)

	binary.LittleEndian.PutUint64(
		hdr[12:20],
		txid,
	)

	binary.LittleEndian.PutUint32(
		hdr[20:24],
		uint32(len(payload)),
	)

	crc :=
		crc32.ChecksumIEEE(
			append(
				hdr[4:24],
				payload...,
			),
		)

	binary.LittleEndian.PutUint32(
		hdr[0:4],
		crc,
	)

	if _, err := w.buf.Write(hdr[:]); err != nil {
		return 0, err
	}

	if _, err := w.buf.Write(payload); err != nil {
		return 0, err
	}

	return w.lsn, nil
}

func (w *WAL) Sync() error {

	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.buf.Flush(); err != nil {
		return err
	}

	return w.file.Sync()
}

type RecordHeader struct {
	CRC32  uint32
	LSN    uint64
	TXID   uint64
	Length uint32
}

func ReadRecord(
	r io.Reader,
) (
	RecordHeader,
	[]byte,
	error,
) {

	var hdr RecordHeader

	var raw [HeaderSize]byte

	_, err := io.ReadFull(
		r,
		raw[:],
	)

	if err != nil {
		return hdr, nil, err
	}

	hdr.CRC32 =
		binary.LittleEndian.Uint32(
			raw[0:4],
		)

	hdr.LSN =
		binary.LittleEndian.Uint64(
			raw[4:12],
		)

	hdr.TXID =
		binary.LittleEndian.Uint64(
			raw[12:20],
		)

	hdr.Length =
		binary.LittleEndian.Uint32(
			raw[20:24],
		)

	payload :=
		make(
			[]byte,
			hdr.Length,
		)

	if hdr.Length > 0 {

		_, err =
			io.ReadFull(
				r,
				payload,
			)

		if err != nil {
			return hdr, nil, err
		}
	}

	return hdr,
		payload,
		nil
}
