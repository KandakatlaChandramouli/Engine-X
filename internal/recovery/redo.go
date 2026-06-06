package recovery

import (
        "encoding/binary"

        "engine-x/core"
)

func ApplyRedo(
        db *core.DB,
        records []RecoveryRecord,
        pageLSN map[uint64]uint64,
) error {

        for _, record := range records {

                payload := record.Payload

                if len(payload) < 8 {
                        continue
                }

                pageID :=
                        binary.LittleEndian.Uint64(
                                payload[:8],
                        )

                current :=
                        pageLSN[pageID]

                if current >= record.Header.LSN {
                        continue
                }

                page :=
                        db.Page(
                                pageID,
                        )

                copy(
                        page,
                        payload[8:],
                )

                pageLSN[pageID] =
                        record.Header.LSN
        }

        return db.Sync()
}
