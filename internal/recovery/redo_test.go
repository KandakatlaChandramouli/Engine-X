package recovery

import (
        "encoding/binary"
        "os"
        "testing"

        "engine-x/core"
        "engine-x/wal"
)

func TestApplyRedo(
        t *testing.T,
) {

        db,
                err :=
                core.Open(
                        "redo.db",
                        4096*4,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("redo.db")
        defer db.Close()

        w,
                err :=
                wal.OpenWAL(
                        "redo.wal",
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("redo.wal")

        payload :=
                make(
                        []byte,
                        4096+8,
                )

        binary.LittleEndian.PutUint64(
                payload[:8],
                0,
        )

        payload[8] = 123

        _, err =
                w.Append(
                        1,
                        payload,
                )

        if err != nil {
                t.Fatal(err)
        }

        if err := w.Sync(); err != nil {
                t.Fatal(err)
        }

        records,
        _,
        err :=
                Scan(
                        "redo.wal",
                )

        if err != nil {
                t.Fatal(err)
        }

        err =
                ApplyRedo(
                        db,
                        records,
                        make(
                                map[uint64]uint64,
                        ),
                )

        if err != nil {
                t.Fatal(err)
        }

        page :=
                db.Page(0)

        if page[0] != 123 {
                t.Fatal()
        }
}
