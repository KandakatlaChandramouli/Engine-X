package aries

type RedoRecord struct {
        LSN    uint64
        PageID uint64
}

func RedoStartLSN(
        dpt *DirtyPageTable,
) uint64 {

        var min uint64

        first := true

        for _, p := range dpt.pages {

                if first || p.RecLSN < min {
                        min = p.RecLSN
                        first = false
                }
        }

        return min
}
