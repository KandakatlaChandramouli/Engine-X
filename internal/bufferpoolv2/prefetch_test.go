package bufferpoolv2

import "testing"

func TestPrefetchAdd(t *testing.T) {
        p := NewPrefetcher()

        p.Add(10)

        if p.Len() != 1 {
                t.Fatal()
        }
}

func TestPrefetchOrder(t *testing.T) {
        p := NewPrefetcher()

        p.Add(1)
        p.Add(2)

        v, ok := p.Next()

        if !ok || v != 1 {
                t.Fatal()
        }
}

func TestPrefetchEmpty(t *testing.T) {
        p := NewPrefetcher()

        _, ok := p.Next()

        if ok {
                t.Fatal()
        }
}
