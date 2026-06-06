package txn

import (
        "sync"
)

type State uint8

const (
        Active State = iota
        Committed
        Aborted
)

type Transaction struct {
        ID       uint64
        State    State
        BeginLSN uint64
        EndLSN   uint64
}

type Manager struct {
        mu     sync.Mutex
        nextID uint64

        active map[uint64]*Transaction
}

func NewManager() *Manager {

        return &Manager{
                nextID: 1,
                active: make(
                        map[uint64]*Transaction,
                ),
        }
}

func (m *Manager) Begin() *Transaction {

        m.mu.Lock()
        defer m.mu.Unlock()

        tx := &Transaction{
                ID:    m.nextID,
                State: Active,
        }

        m.nextID++

        m.active[tx.ID] = tx

        return tx
}

func (m *Manager) Get(
        id uint64,
) (
        *Transaction,
        bool,
) {

        m.mu.Lock()
        defer m.mu.Unlock()

        tx,
                ok :=
                m.active[id]

        return tx,
                ok
}
