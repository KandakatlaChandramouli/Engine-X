package mvcc

import "testing"

func TestSnapshotVisibility(
        t *testing.T,
) {

        tx :=
                BeginTransaction(
                        1,
                )

        v :=
                &Version{
                        BeginTS: tx.StartTS,
                        Value: []byte("x"),
                }

        if !VisibleToTransaction(
                v,
                tx,
        ) {
                t.Fatal()
        }
}

func TestWriteConflict(
        t *testing.T,
) {

        tx1 :=
                BeginTransaction(
                        1,
                )

        tx2 :=
                BeginTransaction(
                        2,
                )

        head :=
                &Version{
                        BeginTS:
                                tx2.StartTS,
                }

        if !HasWriteConflict(
                head,
                tx1,
        ) {
                t.Fatal()
        }
}

func TestCommitTimestamp(
        t *testing.T,
) {

        tx :=
                BeginTransaction(
                        1,
                )

        CommitTransaction(
                tx,
        )

        if tx.CommitTS == 0 {
                t.Fatal()
        }
}
