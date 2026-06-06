package aries

import "testing"

func TestDirtyPageTable(
        t *testing.T,
) {
        dpt := NewDirtyPageTable()

        dpt.Add(
                1,
                100,
        )

        dpt.Add(
                1,
                200,
        )

        e, ok := dpt.Get(1)

        if !ok {
                t.Fatal()
        }

        if e.RecLSN != 100 {
                t.Fatal()
        }

        if dpt.Count() != 1 {
                t.Fatal()
        }
}
