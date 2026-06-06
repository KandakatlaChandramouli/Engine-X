package lockmgr

import "sync"

type Manager struct {
        mu sync.Mutex

        locks map[uint64][]Request
}

func New() *Manager {

        return &Manager{
                locks: make(
                        map[uint64][]Request,
                ),
        }
}
