package bufferpool

import (
        "os"
        "testing"

        "engine-x/core"
)

func TestFetchFromStorage(
        t *testing.T,
) {

        path := "bufferpool.db"

        defer os.Remove(path)

        db,
                err :=
                core.Open(
                        path,
                        4096*4,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer db.Close()

        pool := New(4)

        frame,
                err :=
                pool.FetchFromStorage(
                        db,
                        0,
                )

        if err != nil {
                t.Fatal(err)
        }

        if frame.Data == nil {
                t.Fatal()
        }
}

func TestFlushDirtyPage(
        t *testing.T,
) {

        path := "flush.db"

        defer os.Remove(path)

        db,
                err :=
                core.Open(
                        path,
                        4096*4,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer db.Close()

        pool := New(4)

        frame,
                err :=
                pool.FetchFromStorage(
                        db,
                        0,
                )

        if err != nil {
                t.Fatal(err)
        }

        frame.Data[0] = 99

        if err :=
                pool.MarkDirty(
                        0,
                ); err != nil {
                t.Fatal(err)
        }

        if err :=
                pool.Flush(
                        0,
                ); err != nil {
                t.Fatal(err)
        }

        if frame.Dirty {
                t.Fatal()
        }
}

func TestFlushAll(
        t *testing.T,
) {

        path := "flushall.db"

        defer os.Remove(path)

        db,
                err :=
                core.Open(
                        path,
                        4096*8,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer db.Close()

        pool := New(8)

        for i:=uint64(0); i<4; i++ {

                _,
                err =
                        pool.FetchFromStorage(
                                db,
                                i,
                        )

                if err != nil {
                        t.Fatal(err)
                }

                _ =
                        pool.MarkDirty(
                                i,
                        )
        }

        if err :=
                pool.FlushAll(); err != nil {
                t.Fatal(err)
        }
}
