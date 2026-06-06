package executor

import "testing"

func TestBeginRegistersTransaction(t *testing.T) {
        e := NewEngine()

        e.Begin()

        if e.Table.Count() != 1 {
                t.Fatal()
        }
}

func TestCommitRemovesTransaction(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        e.Commit(tx)

        if e.Table.Count() != 0 {
                t.Fatal()
        }
}

func TestAbortRemovesTransaction(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        e.Abort(tx)

        if e.Table.Count() != 0 {
                t.Fatal()
        }
}
