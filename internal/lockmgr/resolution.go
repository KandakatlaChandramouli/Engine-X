package lockmgr

func (g *WaitsForGraph) YoungestTransaction() uint64 {

        var victim uint64

        for tx := range g.edges {

                if tx > victim {
                        victim = tx
                }
        }

        return victim
}

func (g *WaitsForGraph) ResolveDeadlock() (
        uint64,
        bool,
) {

        if !g.HasCycle() {
                return 0,
                        false
        }

        return g.YoungestTransaction(),
                true
}
