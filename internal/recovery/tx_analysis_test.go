package recovery

import (
        "testing"

        "engine-x/internal/txn"
        "engine-x/wal"
)

func TestAnalyzeTransactions(
        t *testing.T,
) {

        records :=
                []RecoveryRecord{
                        {
                                Header: wal.RecordHeader{
                                        TXID: 1,
                                },
                                Payload: txn.CommitRecord,
                        },
                        {
                                Header: wal.RecordHeader{
                                        TXID: 2,
                                },
                                Payload: []byte("UPDATE"),
                        },
                }

        state :=
                AnalyzeTransactions(
                        records,
                )

        if !state.Committed[1] {
                t.Fatal()
        }

        if !state.Active[2] {
                t.Fatal()
        }
}
