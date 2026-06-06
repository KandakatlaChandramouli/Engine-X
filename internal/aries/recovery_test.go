package aries

import "testing"

func TestRecoveryFlow(
        t *testing.T,
) {
        tt := NewTransactionTable()

        tt.Add(
                TransactionEntry{
                        TxID:   1,
                        Status: Active,
                },
        )

        dpt := NewDirtyPageTable()

        dpt.Add(
                10,
                100,
        )

        if RedoStartLSN(dpt) != 100 {
                t.Fatal()
        }

        undo := BuildUndoList(tt)

        if len(undo) != 1 {
                t.Fatal()
        }

        cp := NewCheckpoint(
                tt,
                dpt,
        )

        if len(cp.Transactions) != 1 {
                t.Fatal()
        }

        if len(cp.DirtyPages) != 1 {
                t.Fatal()
        }
}
