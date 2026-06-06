package aries

type TransactionTable struct {
        entries map[uint64]TransactionEntry
}

func NewTransactionTable() *TransactionTable {

        return &TransactionTable{
                entries: make(
                        map[uint64]TransactionEntry,
                ),
        }
}

func (t *TransactionTable) Add(
        tx TransactionEntry,
) {
        t.entries[tx.TxID] = tx
}

func (t *TransactionTable) Get(
        txID uint64,
) (
        TransactionEntry,
        bool,
) {
        v, ok := t.entries[txID]
        return v, ok
}

func (t *TransactionTable) Count() int {
        return len(t.entries)
}
