package aries

type Checkpoint struct {
        Transactions []TransactionEntry
        DirtyPages   []DirtyPageEntry
}
