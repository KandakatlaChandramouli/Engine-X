package executor

type TransactionTable struct {
        txns map[uint64]*Transaction
}

func NewTransactionTable() *TransactionTable {
        return &TransactionTable{
                txns: make(map[uint64]*Transaction),
        }
}

func (t *TransactionTable) Add(tx *Transaction) {
        t.txns[tx.ID] = tx
}

func (t *TransactionTable) Remove(id uint64) {
        delete(t.txns, id)
}

func (t *TransactionTable) Count() int {
        return len(t.txns)
}
