package mvcc

type Snapshot struct {
        ReadTS uint64
}

func NewSnapshot(
        ts uint64,
) Snapshot {

        return Snapshot{
                ReadTS: ts,
        }
}
