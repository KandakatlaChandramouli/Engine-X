package mvcc

import "sync/atomic"

var globalTS uint64

func NextTS() uint64 {
        return atomic.AddUint64(
                &globalTS,
                1,
        )
}
