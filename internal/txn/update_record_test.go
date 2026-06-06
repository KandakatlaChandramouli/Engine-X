package txn

import (
        "bytes"
        "os"
        "testing"

        "engine-x/wal"
)

func TestUpdateRecordRoundTrip(
        t *testing.T,
) {

        original :=
                UpdateRecord{
                        TxID:   1,
                        PageID: 7,
                        Offset: 32,
                        Before: []byte("old"),
                        After:  []byte("new"),
                }

        encoded :=
                EncodeUpdateRecord(
                        original,
                )

        decoded,
        err :=
                DecodeUpdateRecord(
                        encoded,
                )

        if err != nil {
                t.Fatal(err)
        }

        if decoded.TxID != original.TxID {
                t.Fatal()
        }

        if decoded.PageID != original.PageID {
                t.Fatal()
        }

        if decoded.Offset != original.Offset {
                t.Fatal()
        }

        if !bytes.Equal(
                decoded.Before,
                original.Before,
        ) {
                t.Fatal()
        }

        if !bytes.Equal(
                decoded.After,
                original.After,
        ) {
                t.Fatal()
        }
}

func TestLogUpdate(
        t *testing.T,
) {

        path := "update.wal"

        defer os.Remove(path)

        w,
        err :=
                wal.OpenWAL(
                        path,
                )

        if err != nil {
                t.Fatal(err)
        }

        mgr :=
                NewManager()

        tx :=
                mgr.Begin()

        _, err =
                LogUpdate(
                        tx,
                        w,
                        UpdateRecord{
                                TxID:   tx.ID,
                                PageID: 1,
                                Offset: 8,
                                Before: []byte("a"),
                                After:  []byte("b"),
                        },
                )

        if err != nil {
                t.Fatal(err)
        }
}
