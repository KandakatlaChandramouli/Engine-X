package executor

import "testing"

func TestTransactionTableAdd(t *testing.T) {
        tbl := NewTransactionTable()

        tbl.Add(&Transaction{ID: 1})

        if tbl.Count() != 1 {
                t.Fatal()
        }
}

func TestTransactionTableRemove(t *testing.T) {
        tbl := NewTransactionTable()

        tbl.Add(&Transaction{ID: 1})
        tbl.Remove(1)

        if tbl.Count() != 0 {
                t.Fatal()
        }
}

func TestTransactionTableMultiple(t *testing.T) {
        tbl := NewTransactionTable()

        tbl.Add(&Transaction{ID: 1})
        tbl.Add(&Transaction{ID: 2})
        tbl.Add(&Transaction{ID: 3})

        if tbl.Count() != 3 {
                t.Fatal()
        }
}
