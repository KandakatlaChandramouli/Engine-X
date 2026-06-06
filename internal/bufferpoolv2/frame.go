package bufferpoolv2

type Frame struct {
        PageID   uint64
        Dirty    bool
        PinCount int
}
