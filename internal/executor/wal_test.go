package executor

import "testing"

func TestBeginWritesWAL(t *testing.T) {
        e := NewEngine()

        e.Begin()

        if len(e.WAL.Records) != 1 {
                t.Fatal()
        }

        if e.WAL.Records[0].Type != "BEGIN" {
                t.Fatal()
        }
}

func TestCommitWritesWAL(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        e.Commit(tx)

        if len(e.WAL.Records) != 2 {
                t.Fatal()
        }

        if e.WAL.Records[1].Type != "COMMIT" {
                t.Fatal()
        }
}

func TestAbortWritesWAL(t *testing.T) {
        e := NewEngine()

        tx := e.Begin()

        e.Abort(tx)

        if len(e.WAL.Records) != 2 {
                t.Fatal()
        }

        if e.WAL.Records[1].Type != "ABORT" {
                t.Fatal()
        }
}
