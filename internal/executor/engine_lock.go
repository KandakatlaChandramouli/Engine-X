package executor

func (e *Engine) Acquire(txn uint64, resource string) bool {
        return e.Locks.Acquire(txn, resource)
}

func (e *Engine) Release(txn uint64, resource string) {
        e.Locks.Release(txn, resource)
}

func (e *Engine) IsLocked(resource string) bool {
        return e.Locks.IsLocked(resource)
}
