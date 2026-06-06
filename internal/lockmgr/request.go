package lockmgr

type Request struct {
        TxID       uint64
        ResourceID uint64
        Mode       Mode
        Granted    bool
}
