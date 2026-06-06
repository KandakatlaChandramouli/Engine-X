package concurrency

import "testing"

func TestSearchPath(
        t *testing.T,
) {
        ctx := SearchPath(
                1,
                100,
        )

        if ctx.RootPage != 1 {
                t.Fatal()
        }

        if ctx.Key != 100 {
                t.Fatal()
        }
}
