package latch

import "sync"

type UpgradeableLatch struct {
        mu sync.RWMutex
}

func (l *UpgradeableLatch) RLock() {
        l.mu.RLock()
}

func (l *UpgradeableLatch) RUnlock() {
        l.mu.RUnlock()
}

func (l *UpgradeableLatch) Lock() {
        l.mu.Lock()
}

func (l *UpgradeableLatch) Unlock() {
        l.mu.Unlock()
}

func (l *UpgradeableLatch) Upgrade() {
        l.mu.RUnlock()
        l.mu.Lock()
}

func (l *UpgradeableLatch) Downgrade() {
        l.mu.Unlock()
        l.mu.RLock()
}
