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

func (m *Manager) Acquire(
        txID uint64,
        resourceID uint64,
        mode Mode,
) bool {

        m.mu.Lock()
        defer m.mu.Unlock()

        existing :=
                m.locks[resourceID]

        for _, lock := range existing {

                if lock.Mode == Exclusive {
                        return false
                }

                if mode == Exclusive {
                        return false
                }
        }

        m.locks[resourceID] =
                append(
                        existing,
                        Request{
                                TxID:       txID,
                                ResourceID: resourceID,
                                Mode:       mode,
                                Granted:    true,
                        },
                )

        return true
}

func (m *Manager) Release(
        txID uint64,
        resourceID uint64,
) {

        m.mu.Lock()
        defer m.mu.Unlock()

        locks :=
                m.locks[resourceID]

        filtered :=
                make(
                        []Request,
                        0,
                        len(locks),
                )

        for _, lock := range locks {

                if lock.TxID != txID {
                        filtered =
                                append(
                                        filtered,
                                        lock,
                                )
                }
        }

        if len(filtered) == 0 {
                delete(
                        m.locks,
                        resourceID,
                )
                return
        }

        m.locks[resourceID] =
                filtered
}

func (m *Manager) ReleaseAll(
        txID uint64,
) {

        m.mu.Lock()
        defer m.mu.Unlock()

        for resourceID,
                locks := range m.locks {

                filtered :=
                        make(
                                []Request,
                                0,
                                len(locks),
                        )

                for _, lock := range locks {

                        if lock.TxID != txID {
                                filtered =
                                        append(
                                                filtered,
                                                lock,
                                        )
                        }
                }

                if len(filtered) == 0 {

                        delete(
                                m.locks,
                                resourceID,
                        )

                        continue
                }

                m.locks[resourceID] =
                        filtered
        }
}
