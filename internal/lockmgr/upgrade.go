package lockmgr

func (m *Manager) Upgrade(
        txID uint64,
        resourceID uint64,
) bool {

        m.mu.Lock()
        defer m.mu.Unlock()

        locks :=
                m.locks[resourceID]

        ownerIndex := -1

        for i, lock := range locks {

                if lock.TxID == txID {

                        ownerIndex = i

                        continue
                }

                return false
        }

        if ownerIndex < 0 {
                return false
        }

        locks[ownerIndex].Mode =
                Exclusive

        m.locks[resourceID] =
                locks

        return true
}
