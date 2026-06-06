package latch

import "testing"

func TestCouple(
        t *testing.T,
) {
        c :=
                Couple(
                        1,
                        2,
                )

        if c.Parent != 1 {
                t.Fatal()
        }

        if c.Child != 2 {
                t.Fatal()
        }
}

func TestSafeLeaf(
        t *testing.T,
) {
        if !SafeToRelease(
                true,
                10,
        ) {
                t.Fatal()
        }
}

func TestUnsafeLeaf(
        t *testing.T,
) {
        if SafeToRelease(
                true,
                0,
        ) {
                t.Fatal()
        }
}
