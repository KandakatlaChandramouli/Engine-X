package concurrency

import (
        "sync"
        "testing"
)

func TestConcurrentSearchInsert(
        t *testing.T,
) {
        var wg sync.WaitGroup

        for i := 0; i < 100; i++ {

                wg.Add(1)

                go func(
                        n int,
                ) {
                        defer wg.Done()

                        SearchPath(
                                1,
                                uint64(n),
                        )

                        InsertPath(
                                1,
                                uint64(n),
                                []byte("v"),
                        )
                }(i)
        }

        wg.Wait()
}
