package executor

type Engine struct {
        Manager *TxnManager
        Table   *TransactionTable
        WAL     *WAL
        Locks   *LockManager
}

func NewEngine() *Engine {
        return &Engine{
                Manager: NewTxnManager(),
                Table:   NewTransactionTable(),
                WAL:     NewWAL(),
                Locks:   NewLockManager(),
        }
}

func (e *Engine) Begin() *Transaction {
        tx := e.Manager.Begin()

        e.Table.Add(tx)

        e.WAL.Append(LogRecord{
                TxnID: tx.ID,
                Type:  "BEGIN",
        })

        return tx
}

func (e *Engine) Commit(tx *Transaction) {
        e.Manager.Commit(tx)

        e.WAL.Append(LogRecord{
                TxnID: tx.ID,
                Type:  "COMMIT",
        })

        e.Table.Remove(tx.ID)
}

func (e *Engine) Abort(tx *Transaction) {
        e.Manager.Abort(tx)

        e.WAL.Append(LogRecord{
                TxnID: tx.ID,
                Type:  "ABORT",
        })

        e.Table.Remove(tx.ID)
}
