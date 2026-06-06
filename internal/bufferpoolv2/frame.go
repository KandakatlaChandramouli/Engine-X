package bufferpoolv2

type Frame struct {
        PinCount int
        Dirty    bool
}

func (f *Frame) Pin() {
        f.PinCount++
}

func (f *Frame) Unpin() {
        if f.PinCount > 0 {
                f.PinCount--
        }
}

func (f *Frame) Pinned() bool {
        return f.PinCount > 0
}
