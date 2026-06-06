package executor

type TxnState int

const (
        Active TxnState = iota
        Committed
        Aborted
)

type Transaction struct {
        ID    uint64
        State TxnState
}

type TxnManager struct {
        nextID uint64
}

func NewTxnManager() *TxnManager {
        return &TxnManager{}
}

func (m *TxnManager) Begin() *Transaction {
        m.nextID++

        return &Transaction{
                ID:    m.nextID,
                State: Active,
        }
}
