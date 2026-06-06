package bufferpoolv2

import (
        "sync"
        "testing"
)

func TestConcurrentStats(t *testing.T) {
        s := &Stats{}

        var wg sync.WaitGroup

        for i := 0; i < 100; i++ {
                wg.Add(1)

                go func() {
                        defer wg.Done()
                        s.RecordHit()
                }()
        }

        wg.Wait()

        if s.Hits != 100 {
                t.Fatal()
        }
}

func TestConcurrentPrefetch(t *testing.T) {
        p := NewPrefetcher()

        var wg sync.WaitGroup

        for i := 0; i < 100; i++ {
                wg.Add(1)

                go func(v int) {
                        defer wg.Done()
                        p.Add(v)
                }(i)
        }

        wg.Wait()

        if p.Len() != 100 {
                t.Fatal()
        }
}
