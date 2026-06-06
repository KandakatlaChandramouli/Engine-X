package lockmgr

type WaitsForGraph struct {
        edges map[uint64]map[uint64]struct{}
}

func NewGraph() *WaitsForGraph {

        return &WaitsForGraph{
                edges: make(
                        map[uint64]map[uint64]struct{},
                ),
        }
}

func (g *WaitsForGraph) AddEdge(
        from uint64,
        to uint64,
) {

        if g.edges[from] == nil {

                g.edges[from] =
                        make(
                                map[uint64]struct{},
                        )
        }

        g.edges[from][to] = struct{}{}
}

func (g *WaitsForGraph) RemoveEdge(
        from uint64,
        to uint64,
) {

        if g.edges[from] == nil {
                return
        }

        delete(
                g.edges[from],
                to,
        )
}

func (g *WaitsForGraph) HasCycle() bool {

        visited :=
                make(
                        map[uint64]bool,
                )

        stack :=
                make(
                        map[uint64]bool,
                )

        var dfs func(uint64) bool

        dfs =
                func(node uint64) bool {

                        if stack[node] {
                                return true
                        }

                        if visited[node] {
                                return false
                        }

                        visited[node] = true
                        stack[node] = true

                        for next := range g.edges[node] {

                                if dfs(next) {
                                        return true
                                }
                        }

                        stack[node] = false

                        return false
                }

        for node := range g.edges {

                if dfs(node) {
                        return true
                }
        }

        return false
}
