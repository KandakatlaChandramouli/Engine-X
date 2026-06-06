package aries

type CompensationLogRecord struct {
        TxID        uint64
        UndoLSN     uint64
        NextUndoLSN uint64
}

func NewCLR(
        txID uint64,
        undoLSN uint64,
        nextUndoLSN uint64,
) CompensationLogRecord {

        return CompensationLogRecord{
                TxID:        txID,
                UndoLSN:     undoLSN,
                NextUndoLSN: nextUndoLSN,
        }
}
