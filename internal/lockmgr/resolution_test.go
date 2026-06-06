package lockmgr

import "testing"

func TestResolveDeadlock(
        t *testing.T,
) {

        g := NewGraph()

        g.AddEdge(
                1,
                2,
        )

        g.AddEdge(
                2,
                3,
        )

        g.AddEdge(
                3,
                1,
        )

        victim,
                ok :=
                g.ResolveDeadlock()

        if !ok {
                t.Fatal()
        }

        if victim != 3 {
                t.Fatal()
        }
}

func TestResolveNoDeadlock(
        t *testing.T,
) {

        g := NewGraph()

        g.AddEdge(
                1,
                2,
        )

        victim,
                ok :=
                g.ResolveDeadlock()

        if ok {
                t.Fatal()
        }

        if victim != 0 {
                t.Fatal()
        }
}
