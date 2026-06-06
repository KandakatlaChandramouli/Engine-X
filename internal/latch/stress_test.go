package latch

import (
        "sync"
        "testing"
)

func TestConcurrentReaders(
        t *testing.T,
) {
        l := &Latch{}

        var wg sync.WaitGroup

        for i := 0; i < 100; i++ {

                wg.Add(1)

                go func() {
                        defer wg.Done()

                        l.RLock()
                        l.RUnlock()
                }()
        }

        wg.Wait()
}
