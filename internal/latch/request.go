package latch

func (l *Latch) RLock() {
        l.mu.RLock()
}

func (l *Latch) RUnlock() {
        l.mu.RUnlock()
}

func (l *Latch) Lock() {
        l.mu.Lock()
}

func (l *Latch) Unlock() {
        l.mu.Unlock()
}
