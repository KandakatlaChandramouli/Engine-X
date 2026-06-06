package aries

import "testing"

func TestRedoStartLSN(
        t *testing.T,
) {
        dpt := NewDirtyPageTable()

        dpt.Add(
                1,
                300,
        )

        dpt.Add(
                2,
                100,
        )

        dpt.Add(
                3,
                200,
        )

        if RedoStartLSN(dpt) != 100 {
                t.Fatal()
        }
}
