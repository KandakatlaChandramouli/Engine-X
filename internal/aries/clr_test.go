package aries

import "testing"

func TestCLR(
        t *testing.T,
) {
        clr := NewCLR(
                1,
                100,
                50,
        )

        if clr.TxID != 1 {
                t.Fatal()
        }

        if clr.UndoLSN != 100 {
                t.Fatal()
        }

        if clr.NextUndoLSN != 50 {
                t.Fatal()
        }
}
