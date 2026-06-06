package txn

import (
        "engine-x/wal"
)

func LogUpdate(
        tx *Transaction,
        w *wal.WAL,
        record UpdateRecord,
) (
        uint64,
        error,
) {

        payload :=
                EncodeUpdateRecord(
                        record,
                )

        return w.Append(
                tx.ID,
                payload,
        )
}
