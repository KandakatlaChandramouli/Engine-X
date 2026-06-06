package latch

import "testing"

func TestSearchPath(
        t *testing.T,
) {
        var p SearchPath

        p.Add(1)
        p.Add(2)
        p.Add(3)

        if p.Depth() != 3 {
                t.Fatal()
        }
}
