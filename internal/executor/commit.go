package executor

func (m *TxnManager) Commit(tx *Transaction) {
        tx.State = Committed
}

func (m *TxnManager) Abort(tx *Transaction) {
        tx.State = Aborted
}
