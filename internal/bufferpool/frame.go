package bufferpool

type Frame struct {
        PageID uint64
        Dirty  bool
        Pin    uint32
}
