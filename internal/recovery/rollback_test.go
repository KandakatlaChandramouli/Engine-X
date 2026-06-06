package recovery

import (
        "os"
        "testing"

        "engine-x/core"
        "engine-x/internal/txn"
)

func TestRollbackTransaction(
        t *testing.T,
) {

        db,
        err :=
                core.Open(
                        "rollback.db",
                        4096,
                )

        if err != nil {
                t.Fatal(err)
        }

        defer os.Remove("rollback.db")
        defer db.Close()

        page :=
                db.Page(0)

        page[100] = 'B'

        updates :=
                []txn.UpdateRecord{
                        {
                                TxID:   1,
                                PageID: 0,
                                Offset: 100,
                                Before: []byte{'A'},
                                After:  []byte{'B'},
                        },
                }

        err =
                RollbackTransaction(
                        db,
                        updates,
                )

        if err != nil {
                t.Fatal(err)
        }

        if page[100] != 'A' {
                t.Fatal()
        }
}
