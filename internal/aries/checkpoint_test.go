package aries

import "testing"

func TestCheckpoint(
        t *testing.T,
) {
        tt := NewTransactionTable()

        tt.Add(
                TransactionEntry{
                        TxID: 1,
                },
        )

        dpt := NewDirtyPageTable()

        dpt.Add(
                10,
                100,
        )

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
