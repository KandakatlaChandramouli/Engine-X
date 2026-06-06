package latch

import "testing"

func TestUpgrade(
        t *testing.T,
) {
        l := &UpgradeableLatch{}

        l.RLock()
        l.Upgrade()
        l.Unlock()
}

func TestDowngrade(
        t *testing.T,
) {
        l := &UpgradeableLatch{}

        l.Lock()
        l.Downgrade()
        l.RUnlock()
}
