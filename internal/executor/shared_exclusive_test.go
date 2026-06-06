package executor

import "testing"

func TestSharedSharedCompatible(t *testing.T) {
        lm := NewLockManager()

        if !lm.AcquireShared(1, "page1") {
                t.Fatal()
        }

        if !lm.AcquireShared(2, "page1") {
                t.Fatal()
        }
}

func TestExclusiveBlocksShared(t *testing.T) {
        lm := NewLockManager()

        lm.AcquireExclusive(1, "page1")

        if lm.AcquireShared(2, "page1") {
                t.Fatal()
        }
}

func TestSharedBlocksExclusive(t *testing.T) {
        lm := NewLockManager()

        lm.AcquireShared(1, "page1")

        if lm.AcquireExclusive(2, "page1") {
                t.Fatal()
        }
}
