package core

import (
	"errors"
	"os"

	"golang.org/x/sys/unix"
)

type DB struct {
	file *os.File
	mmap []byte
	size int64
}

func Open(
	path string,
	size int64,
) (*DB, error) {

	if size <= 0 {
		return nil, errors.New("invalid size")
	}

	f, err :=
		os.OpenFile(
			path,
			os.O_CREATE|os.O_RDWR,
			0644,
		)

	if err != nil {
		return nil, err
	}

	if err := f.Truncate(size); err != nil {
		_ = f.Close()
		return nil, err
	}

	data, err :=
		unix.Mmap(
			int(f.Fd()),
			0,
			int(size),
			unix.PROT_READ|
				unix.PROT_WRITE,
			unix.MAP_SHARED,
		)

	if err != nil {
		_ = f.Close()
		return nil, err
	}

	return &DB{
		file: f,
		mmap: data,
		size: size,
	}, nil
}

func (db *DB) Page(id uint64) []byte {

	start :=
		int(id) * PageSize

	end :=
		start + PageSize

	return db.mmap[start:end]
}

func (db *DB) PageCount() uint64 {
	return uint64(db.size) / PageSize
}

func (db *DB) Sync() error {
	return unix.Msync(
		db.mmap,
		unix.MS_SYNC,
	)
}

func (db *DB) Close() error {

	if err := unix.Munmap(
		db.mmap,
	); err != nil {
		return err
	}

	return db.file.Close()
}
