package aries

type UndoRecord struct {
        TxID uint64
        LSN  uint64
}

func BuildUndoList(
        tt *TransactionTable,
) []uint64 {

        var undo []uint64

        for _, tx := range tt.entries {

                if tx.Status == Active {
                        undo = append(
                                undo,
                                tx.TxID,
                        )
                }
        }

        return undo
}
