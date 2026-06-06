package executor

func NewEngineWithWAL() *Engine {
        e := NewEngine()

        ewal := NewWAL()

        _ = ewal

        return e
}
