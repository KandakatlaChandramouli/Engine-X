package executor

type LockMode int

const (
        Shared LockMode = iota
        Exclusive
)

type LockState struct {
        Mode    LockMode
        Holders map[uint64]bool
}

type LockManager struct {
        locks map[string]*LockState
}

func NewLockManager() *LockManager {
        return &LockManager{
                locks: make(map[string]*LockState),
        }
}

func (l *LockManager) AcquireShared(txn uint64, resource string) bool {
        state, ok := l.locks[resource]

        if !ok {
                l.locks[resource] = &LockState{
                        Mode: Shared,
                        Holders: map[uint64]bool{
                                txn: true,
                        },
                }
                return true
        }

        if state.Mode == Exclusive {
                return false
        }

        state.Holders[txn] = true
        return true
}

func (l *LockManager) AcquireExclusive(txn uint64, resource string) bool {
        _, ok := l.locks[resource]

        if ok {
                return false
        }

        l.locks[resource] = &LockState{
                Mode: Exclusive,
                Holders: map[uint64]bool{
                        txn: true,
                },
        }

        return true
}

func (l *LockManager) Release(txn uint64, resource string) {
        state, ok := l.locks[resource]

        if !ok {
                return
        }

        delete(state.Holders, txn)

        if len(state.Holders) == 0 {
                delete(l.locks, resource)
        }
}

func (l *LockManager) IsLocked(resource string) bool {
        _, ok := l.locks[resource]
        return ok
}
