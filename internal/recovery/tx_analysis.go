package recovery

import (
        "bytes"

        "engine-x/internal/txn"
)

type RecoveryState struct {
        Committed map[uint64]bool
        Aborted   map[uint64]bool
        Active    map[uint64]bool
}

func AnalyzeTransactions(
        records []RecoveryRecord,
) RecoveryState {

        state :=
                RecoveryState{
                        Committed: make(map[uint64]bool),
                        Aborted:   make(map[uint64]bool),
                        Active:    make(map[uint64]bool),
                }

        for _, r := range records {

                txid :=
                        r.Header.TXID

                state.Active[txid] = true

                if bytes.Equal(
                        r.Payload,
                        txn.CommitRecord,
                ) {

                        state.Committed[txid] = true
                        delete(
                                state.Active,
                                txid,
                        )
                }

                if bytes.Equal(
                        r.Payload,
                        txn.AbortRecord,
                ) {

                        state.Aborted[txid] = true
                        delete(
                                state.Active,
                                txid,
                        )
                }
        }

        return state
}
