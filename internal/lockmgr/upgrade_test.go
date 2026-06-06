package lockmgr

import "testing"

func TestUpgradeSuccess(
        t *testing.T,
) {

        lm := New()

        if !lm.Acquire(
                1,
                100,
                Shared,
        ) {
                t.Fatal()
        }

        if !lm.Upgrade(
                1,
                100,
        ) {
                t.Fatal()
        }
}

func TestUpgradeConflict(
        t *testing.T,
) {

        lm := New()

        lm.Acquire(
                1,
                100,
                Shared,
        )

        lm.Acquire(
                2,
                100,
                Shared,
        )

        if lm.Upgrade(
                1,
                100,
        ) {
                t.Fatal()
        }
}

func TestUpgradeMissingLock(
        t *testing.T,
) {

        lm := New()

        if lm.Upgrade(
                1,
                100,
        ) {
                t.Fatal()
        }
}
