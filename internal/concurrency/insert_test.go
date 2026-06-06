package concurrency

import "testing"

func TestInsertPath(
        t *testing.T,
) {
        ctx := InsertPath(
                1,
                100,
                []byte("value"),
        )

        if ctx.RootPage != 1 {
                t.Fatal()
        }

        if ctx.Key != 100 {
                t.Fatal()
        }

        if string(ctx.Value) != "value" {
                t.Fatal()
        }
}
