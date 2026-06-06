package executor

type Engine struct {
        Manager *TxnManager
        Table   *TransactionTable
}

func NewEngine() *Engine {
        return &Engine{
                Manager: NewTxnManager(),
                Table:   NewTransactionTable(),
        }
}

func (e *Engine) Begin() *Transaction {
        tx := e.Manager.Begin()
        e.Table.Add(tx)
        return tx
}

func (e *Engine) Commit(tx *Transaction) {
        e.Manager.Commit(tx)
        e.Table.Remove(tx.ID)
}

func (e *Engine) Abort(tx *Transaction) {
        e.Manager.Abort(tx)
        e.Table.Remove(tx.ID)
}
