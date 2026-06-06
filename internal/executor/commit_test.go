package executor

import "testing"

func TestCommit(t *testing.T) {
        mgr := NewTxnManager()

        tx := mgr.Begin()

        mgr.Commit(tx)

        if tx.State != Committed {
                t.Fatal()
        }
}

func TestAbort(t *testing.T) {
        mgr := NewTxnManager()

        tx := mgr.Begin()

        mgr.Abort(tx)

        if tx.State != Aborted {
                t.Fatal()
        }
}

func TestCommitDoesNotChangeID(t *testing.T) {
        mgr := NewTxnManager()

        tx := mgr.Begin()

        id := tx.ID

        mgr.Commit(tx)

        if tx.ID != id {
                t.Fatal()
        }
}
