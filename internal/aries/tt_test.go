package aries

import "testing"

func TestTransactionTable(
        t *testing.T,
) {
        tt := NewTransactionTable()

        tt.Add(
                TransactionEntry{
                        TxID:    1,
                        Status:  Active,
                        LastLSN: 100,
                },
        )

        tx, ok := tt.Get(1)

        if !ok {
                t.Fatal()
        }

        if tx.LastLSN != 100 {
                t.Fatal()
        }

        if tt.Count() != 1 {
                t.Fatal()
        }
}
