package aries

type Checkpoint struct {
        Transactions []TransactionEntry
        DirtyPages    []DirtyPageEntry
}

func NewCheckpoint(
        tt *TransactionTable,
        dpt *DirtyPageTable,
) Checkpoint {

        cp := Checkpoint{}

        for _, tx := range tt.entries {
                cp.Transactions = append(
                        cp.Transactions,
                        tx,
                )
        }

        for _, dp := range dpt.pages {
                cp.DirtyPages = append(
                        cp.DirtyPages,
                        dp,
                )
        }

        return cp
}
