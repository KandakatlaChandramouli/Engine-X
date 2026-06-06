package lockmgr

import "testing"

func TestGraphNoCycle(
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

        if g.HasCycle() {
                t.Fatal()
        }
}

func TestGraphCycle(
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

        if !g.HasCycle() {
                t.Fatal()
        }
}

func TestRemoveEdge(
        t *testing.T,
) {

        g := NewGraph()

        g.AddEdge(
                1,
                2,
        )

        g.AddEdge(
                2,
                1,
        )

        if !g.HasCycle() {
                t.Fatal()
        }

        g.RemoveEdge(
                2,
                1,
        )

        if g.HasCycle() {
                t.Fatal()
        }
}
