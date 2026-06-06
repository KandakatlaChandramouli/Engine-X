package concurrency

import "testing"

func TestAcquireRoot(
        t *testing.T,
) {
        root := AcquireRoot(
                1,
        )

        if root.RootPage != 1 {
                t.Fatal()
        }

        if !root.Held {
                t.Fatal()
        }
}
