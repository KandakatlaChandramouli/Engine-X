package wal

import (
        "bytes"
        "os"
        "testing"
)

func TestOpenWAL(
        t *testing.T,
) {
        path := "test.wal"

        defer os.Remove(path)

        w, err := OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        if err := w.file.Close(); err != nil {
                t.Fatal(err)
        }
}

func TestAppendAndSync(
        t *testing.T,
) {
        path := "append.wal"

        defer os.Remove(path)

        w, err := OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        lsn, err :=
                w.Append(
                        42,
                        []byte("payload"),
                )

        if err != nil {
                t.Fatal(err)
        }

        if lsn != 1 {
                t.Fatalf(
                        "expected lsn 1 got %d",
                        lsn,
                )
        }

        if err := w.Sync(); err != nil {
                t.Fatal(err)
        }

        _ = w.file.Close()
}

func TestReadRecord(
        t *testing.T,
) {
        path := "record.wal"

        defer os.Remove(path)

        w, err := OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        _, err =
                w.Append(
                        99,
                        []byte("hello"),
                )

        if err != nil {
                t.Fatal(err)
        }

        if err := w.Sync(); err != nil {
                t.Fatal(err)
        }

        _ = w.file.Close()

        f, err := os.Open(path)

        if err != nil {
                t.Fatal(err)
        }

        defer f.Close()

        hdr,
        payload,
        err :=
                ReadRecord(f)

        if err != nil {
                t.Fatal(err)
        }

        if hdr.LSN != 1 {
                t.Fatal()
        }

        if hdr.TXID != 99 {
                t.Fatal()
        }

        if !bytes.Equal(
                payload,
                []byte("hello"),
        ) {
                t.Fatal()
        }
}

func TestAppendSequence(
        t *testing.T,
) {
        path := "sequence.wal"

        defer os.Remove(path)

        w, err := OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        for i := 1; i <= 100; i++ {

                lsn, err :=
                        w.Append(
                                uint64(i),
                                []byte("x"),
                        )

                if err != nil {
                        t.Fatal(err)
                }

                if lsn != uint64(i) {
                        t.Fatal()
                }
        }

        if err := w.Sync(); err != nil {
                t.Fatal(err)
        }

        _ = w.file.Close()
}

func TestReadRecordEOF(
        t *testing.T,
) {
        _, _, err :=
                ReadRecord(
                        bytes.NewReader(nil),
                )

        if err == nil {
                t.Fatal()
        }
}
