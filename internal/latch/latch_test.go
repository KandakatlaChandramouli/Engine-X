package latch

import "testing"

func TestReadLatch(
        t *testing.T,
) {
        l := &Latch{}

        l.RLock()
        l.RUnlock()
}

func TestWriteLatch(
        t *testing.T,
) {
        l := &Latch{}

        l.Lock()
        l.Unlock()
}

func TestManager(
        t *testing.T,
) {
        m := NewManager()

        a := m.Get(1)
        b := m.Get(1)

        if a != b {
                t.Fatal()
        }
}
