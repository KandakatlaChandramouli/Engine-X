package aries

type TransactionEntry struct {
        TxID uint64
}

type DirtyPageEntry struct {
        PageID uint64
        RecLSN uint64
}
