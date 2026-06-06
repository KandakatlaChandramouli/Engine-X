package lockmgr

import "testing"

func TestAcquireSharedShared(
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

        if !lm.Acquire(
                2,
                100,
                Shared,
        ) {
                t.Fatal()
        }
}

func TestAcquireSharedExclusiveDenied(
        t *testing.T,
) {

        lm := New()

        lm.Acquire(
                1,
                100,
                Shared,
        )

        if lm.Acquire(
                2,
                100,
                Exclusive,
        ) {
                t.Fatal()
        }
}

func TestAcquireExclusiveSharedDenied(
        t *testing.T,
) {

        lm := New()

        lm.Acquire(
                1,
                100,
                Exclusive,
        )

        if lm.Acquire(
                2,
                100,
                Shared,
        ) {
                t.Fatal()
        }
}

func TestRelease(
        t *testing.T,
) {

        lm := New()

        lm.Acquire(
                1,
                100,
                Shared,
        )

        lm.Release(
                1,
                100,
        )

        if !lm.Acquire(
                2,
                100,
                Exclusive,
        ) {
                t.Fatal()
        }
}

func TestReleaseAll(
        t *testing.T,
) {

        lm := New()

        lm.Acquire(
                1,
                100,
                Shared,
        )

        lm.Acquire(
                1,
                200,
                Shared,
        )

        lm.ReleaseAll(
                1,
        )

        if !lm.Acquire(
                2,
                100,
                Exclusive,
        ) {
                t.Fatal()
        }

        if !lm.Acquire(
                2,
                200,
                Exclusive,
        ) {
                t.Fatal()
        }
}
