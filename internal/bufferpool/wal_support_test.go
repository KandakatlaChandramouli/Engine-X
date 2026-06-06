package bufferpool

import (
        "os"
        "testing"

        "engine-x/core"
        "engine-x/wal"
)

func TestWritePageLog(
        t *testing.T,
) {

        db,
                err :=
                core.Open(
                        "wal_page.db",
                        4096,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("wal_page.db")
        defer db.Close()

        w,
                err :=
                wal.OpenWAL(
                        "wal_page.log",
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("wal_page.log")
        defer w.Sync()

        pool := New(2)

        _,
        err =
                pool.FetchFromStorage(
                        db,
                        0,
                )

        if err != nil {
                t.Fatal(err)
        }

        err =
                pool.WritePageLog(
                        0,
                        1,
                        []byte("update"),
                        w,
                )

        if err != nil {
                t.Fatal(err)
        }

        idx :=
                pool.PageTable[0]

        if pool.Frames[idx].LSN == 0 {
                t.Fatal()
        }
}

func TestFlushWithWAL(
        t *testing.T,
) {

        db,
                err :=
                core.Open(
                        "flushwal.db",
                        4096,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("flushwal.db")
        defer db.Close()

        w,
                err :=
                wal.OpenWAL(
                        "flushwal.log",
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("flushwal.log")

        pool := New(2)

        frame,
                err :=
                pool.FetchFromStorage(
                        db,
                        0,
                )

        if err != nil {
                t.Fatal(err)
        }

        frame.Data[0] = 7

        _ =
                pool.MarkDirty(
                        0,
                )

        err =
                pool.WritePageLog(
                        0,
                        1,
                        []byte("change"),
                        w,
                )

        if err != nil {
                t.Fatal(err)
        }

        err =
                pool.FlushWithWAL(
                        0,
                        w,
                )

        if err != nil {
                t.Fatal(err)
        }
}
