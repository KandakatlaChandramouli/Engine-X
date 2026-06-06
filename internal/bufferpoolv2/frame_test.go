package bufferpoolv2

import "testing"

func TestPin(t *testing.T) {
        f := &Frame{}

        f.Pin()

        if f.PinCount != 1 {
                t.Fatal()
        }
}

func TestUnpin(t *testing.T) {
        f := &Frame{
                PinCount: 1,
        }

        f.Unpin()

        if f.PinCount != 0 {
                t.Fatal()
        }
}

func TestPinned(t *testing.T) {
        f := &Frame{}

        if f.Pinned() {
                t.Fatal()
        }

        f.Pin()

        if !f.Pinned() {
                t.Fatal()
        }
}
