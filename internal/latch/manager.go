package latch

import "sync"

type Manager struct {
        mu      sync.Mutex
        latches map[uint64]*Latch
}

func NewManager() *Manager {
        return &Manager{
                latches: make(
                        map[uint64]*Latch,
                ),
        }
}

func (m *Manager) Get(
        pageID uint64,
) *Latch {

        m.mu.Lock()
        defer m.mu.Unlock()

        l, ok := m.latches[pageID]

        if ok {
                return l
        }

        l = &Latch{}

        m.latches[pageID] = l

        return l
}
