package bufferpool

import "engine-x/core"

type Frame struct {
        PageID uint64
        Dirty  bool
        Pin    uint32
        Used   bool

        Data []byte

        DB *core.DB
}
