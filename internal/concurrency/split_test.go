package concurrency

import "testing"

func TestCanSplit(
        t *testing.T,
) {
        ctx := CanSplit(
                10,
                true,
        )

        if ctx.PageID != 10 {
                t.Fatal()
        }

        if !ctx.Safe {
                t.Fatal()
        }
}
