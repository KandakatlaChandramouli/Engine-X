package executor

import "testing"

func TestAcquireLock(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        if !e.Acquire(tx.ID, "page1") {
                t.Fatal()
        }
}

func TestReleaseLock(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        e.Acquire(tx.ID, "page1")
        e.Release(tx.ID, "page1")

        if e.IsLocked("page1") {
                t.Fatal()
        }
}
