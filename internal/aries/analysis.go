package aries

type TransactionStatus uint8

const (
        Active TransactionStatus = iota
        Committed
        Aborted
)

type TransactionEntry struct {
        TxID   uint64
        Status TransactionStatus
        LastLSN uint64
}

type DirtyPageEntry struct {
        PageID uint64
        RecLSN uint64
}

type AnalysisResult struct {
        Transactions map[uint64]TransactionEntry
        DirtyPages   map[uint64]DirtyPageEntry
}

func Analyze(
        txns []TransactionEntry,
        pages []DirtyPageEntry,
) AnalysisResult {

        r := AnalysisResult{
                Transactions: make(map[uint64]TransactionEntry),
                DirtyPages:   make(map[uint64]DirtyPageEntry),
        }

        for _, tx := range txns {
                r.Transactions[tx.TxID] = tx
        }

        for _, p := range pages {
                r.DirtyPages[p.PageID] = p
        }

        return r
}
