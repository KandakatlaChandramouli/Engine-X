package bufferpoolv2

import "testing"

func TestDirtyPage(t *testing.T) {
        f := &Frame{}

        f.Dirty = true

        if !f.Dirty {
                t.Fatal()
        }
}

func TestCleanPage(t *testing.T) {
        f := &Frame{
                Dirty: true,
        }

        f.Dirty = false

        if f.Dirty {
                t.Fatal()
        }
}

func TestEvictable(t *testing.T) {
        f := &Frame{}

        if !f.Evictable() {
                t.Fatal()
        }

        f.Pin()

        if f.Evictable() {
                t.Fatal()
        }
}
