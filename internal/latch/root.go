package latch

import "sync"

type RootLatch struct {
        mu sync.RWMutex
}

func (r *RootLatch) RLock() {
        r.mu.RLock()
}

func (r *RootLatch) RUnlock() {
        r.mu.RUnlock()
}

func (r *RootLatch) Lock() {
        r.mu.Lock()
}

func (r *RootLatch) Unlock() {
        r.mu.Unlock()
}
