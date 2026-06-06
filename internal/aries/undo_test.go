package aries

import "testing"

func TestBuildUndoList(
        t *testing.T,
) {
        tt := NewTransactionTable()

        tt.Add(
                TransactionEntry{
                        TxID:   1,
                        Status: Active,
                },
        )

        tt.Add(
                TransactionEntry{
                        TxID:   2,
                        Status: Committed,
                },
        )

        undo := BuildUndoList(tt)

        if len(undo) != 1 {
                t.Fatal()
        }

        if undo[0] != 1 {
                t.Fatal()
        }
}
