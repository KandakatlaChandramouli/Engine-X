package latch

import "testing"

func TestRootLatch(
        t *testing.T,
) {
        var r RootLatch

        r.RLock()
        r.RUnlock()

        r.Lock()
        r.Unlock()
}
