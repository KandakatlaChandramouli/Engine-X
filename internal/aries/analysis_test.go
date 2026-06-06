package aries

import "testing"

func TestAnalyze(
        t *testing.T,
) {
        r := Analyze(
                []TransactionEntry{
                        {
                                TxID: 1,
                                Status: Active,
                        },
                },
                []DirtyPageEntry{
                        {
                                PageID: 10,
                                RecLSN: 100,
                        },
                },
        )

        if len(r.Transactions) != 1 {
                t.Fatal()
        }

        if len(r.DirtyPages) != 1 {
                t.Fatal()
        }
}
