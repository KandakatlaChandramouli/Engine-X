package txn

import (
        "os"
        "testing"

        "engine-x/wal"
)

func TestBeginTransaction(
        t *testing.T,
) {

        mgr := NewManager()

        tx :=
                mgr.Begin()

        if tx.ID == 0 {
                t.Fatal()
        }

        if tx.State != Active {
                t.Fatal()
        }
}

func TestCommitTransaction(
        t *testing.T,
) {

        path := "commit.wal"

        defer os.Remove(path)

        w,
                err :=
                wal.OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        mgr := NewManager()

        tx :=
                mgr.Begin()

        err =
                mgr.Commit(
                        tx,
                        w,
                )

        if err != nil {
                t.Fatal(err)
        }

        if tx.State != Committed {
                t.Fatal()
        }
}

func TestAbortTransaction(
        t *testing.T,
) {

        path := "abort.wal"

        defer os.Remove(path)

        w,
                err :=
                wal.OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        mgr := NewManager()

        tx :=
                mgr.Begin()

        err =
                mgr.Abort(
                        tx,
                        w,
                )

        if err != nil {
                t.Fatal(err)
        }

        if tx.State != Aborted {
                t.Fatal()
        }
}
