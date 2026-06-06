package lockmgr

func SerializableRead(
        lm *Manager,
        txID uint64,
        resourceID uint64,
) bool {

        return lm.Acquire(
                txID,
                resourceID,
                Shared,
        )
}

func SerializableWrite(
        lm *Manager,
        txID uint64,
        resourceID uint64,
) bool {

        return lm.Acquire(
                txID,
                resourceID,
                Exclusive,
        )
}
