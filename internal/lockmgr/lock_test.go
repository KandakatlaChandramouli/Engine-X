package lockmgr

import "testing"

func TestSharedShared(
        t *testing.T,
) {
        if !Compatible(
                Shared,
                Shared,
        ) {
                t.Fatal()
        }
}

func TestSharedExclusive(
        t *testing.T,
) {
        if Compatible(
                Shared,
                Exclusive,
        ) {
                t.Fatal()
        }
}

func TestExclusiveShared(
        t *testing.T,
) {
        if Compatible(
                Exclusive,
                Shared,
        ) {
                t.Fatal()
        }
}
