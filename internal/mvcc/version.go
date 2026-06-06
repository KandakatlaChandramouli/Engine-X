package mvcc

type Version struct {
        TxID    uint64
        BeginTS uint64
        EndTS   uint64

        Deleted bool

        Value []byte

        Prev *Version
}

func Visible(
        v *Version,
        readTS uint64,
) bool {

        if v == nil {
                return false
        }

        if v.BeginTS > readTS {
                return false
        }

        if v.EndTS != 0 &&
                readTS >= v.EndTS {

                return false
        }

        if v.Deleted {
                return false
        }

        return true
}
