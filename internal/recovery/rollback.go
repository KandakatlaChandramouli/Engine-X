package recovery

import (
        "engine-x/core"
        "engine-x/internal/txn"
)

func RollbackTransaction(
        db *core.DB,
        updates []txn.UpdateRecord,
) error {

        for i := len(updates) - 1; i >= 0; i-- {

                record := updates[i]

                page :=
                        db.Page(
                                record.PageID,
                        )

                start :=
                        int(record.Offset)

                end :=
                        start +
                        len(record.Before)

                copy(
                        page[start:end],
                        record.Before,
                )
        }

        return db.Sync()
}
