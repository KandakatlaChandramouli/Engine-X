package bufferpoolv2

func Flush(f *Frame) {
        f.Dirty = false
}

func RunFlusher(frames []*Frame) {
        for _, f := range frames {
                if f.Dirty {
                        Flush(f)
                }
        }
}
