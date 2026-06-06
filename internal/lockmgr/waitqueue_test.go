package lockmgr

import "testing"

func TestQueueFIFO(
        t *testing.T,
) {

        var q WaitQueue

        q.Enqueue(
                Request{
                        TxID: 1,
                },
        )

        q.Enqueue(
                Request{
                        TxID: 2,
                },
        )

        r,
                ok :=
                q.Dequeue()

        if !ok ||
                r.TxID != 1 {
                t.Fatal()
        }

        r,
                ok =
                q.Dequeue()

        if !ok ||
                r.TxID != 2 {
                t.Fatal()
        }
}

func TestQueueEmpty(
        t *testing.T,
) {

        var q WaitQueue

        _,
                ok :=
                q.Dequeue()

        if ok {
                t.Fatal()
        }
}

func TestQueueLength(
        t *testing.T,
) {

        var q WaitQueue

        q.Enqueue(
                Request{
                        TxID: 1,
                },
        )

        q.Enqueue(
                Request{
                        TxID: 2,
                },
        )

        if q.Len() != 2 {
                t.Fatal()
        }
}
