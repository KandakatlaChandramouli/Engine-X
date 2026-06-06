package executor

type LogRecord struct {
        TxnID uint64
        Type  string
}

type WAL struct {
        Records []LogRecord
}

func NewWAL() *WAL {
        return &WAL{}
}

func (w *WAL) Append(r LogRecord) {
        w.Records = append(w.Records, r)
}
