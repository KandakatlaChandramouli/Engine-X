package latch

import "testing"

func TestNeedSplit(
        t *testing.T,
) {
        if !NeedSplit(0) {
                t.Fatal()
        }

        if NeedSplit(10) {
                t.Fatal()
        }
}
