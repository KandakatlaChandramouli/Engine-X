package latch

import "testing"

func TestCrabbingPath(
        t *testing.T,
) {
        var p CrabbingPath

        p.Push(1)
        p.Push(2)

        if p.Pop() != 2 {
                t.Fatal()
        }

        if p.Pop() != 1 {
                t.Fatal()
        }
}
