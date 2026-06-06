package bufferpoolv2

import "testing"

func TestFlushDirty(t *testing.T) {
        f := &Frame{
                Dirty: true,
        }

        Flush(f)

        if f.Dirty {
                t.Fatal()
        }
}

func TestFlushClean(t *testing.T) {
        f := &Frame{}

        Flush(f)

        if f.Dirty {
                t.Fatal()
        }
}

func TestBackgroundFlusher(t *testing.T) {
        f := &Frame{
                Dirty: true,
        }

        RunFlusher([]*Frame{f})

        if f.Dirty {
                t.Fatal()
        }
}
