package storage

type RecoveryHook struct {
	dpt *DirtyPageTable
}

func NewRecoveryHook(dpt *DirtyPageTable) *RecoveryHook {
	return &RecoveryHook{
		dpt: dpt,
	}
}

func (r *RecoveryHook) RecoveryStartLSN() uint64 {
	var min uint64

	first := true

	for _, lsn := range r.dpt.pages {
		if first || lsn < min {
			min = lsn
			first = false
		}
	}

	return min
}
