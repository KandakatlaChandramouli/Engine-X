package executor

import "testing"

func TestBeginTransaction(t *testing.T) {
        mgr := NewTxnManager()

        tx := mgr.Begin()

        if tx.ID != 1 {
                t.Fatal()
        }

        if tx.State != Active {
                t.Fatal()
        }
}

func TestMultipleTransactions(t *testing.T) {
        mgr := NewTxnManager()

        tx1 := mgr.Begin()
        tx2 := mgr.Begin()

        if tx1.ID != 1 {
                t.Fatal()
        }

        if tx2.ID != 2 {
                t.Fatal()
        }
}
