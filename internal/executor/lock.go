package executor

type LockManager struct {
        locks map[string]uint64
}

func NewLockManager() *LockManager {
        return &LockManager{
                locks: make(map[string]uint64),
        }
}

func (l *LockManager) Acquire(txn uint64, resource string) bool {
        if _, ok := l.locks[resource]; ok {
                return false
        }

        l.locks[resource] = txn
        return true
}

func (l *LockManager) Release(txn uint64, resource string) {
        delete(l.locks, resource)
}

func (l *LockManager) IsLocked(resource string) bool {
        _, ok := l.locks[resource]
        return ok
}
