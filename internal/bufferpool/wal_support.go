package bufferpool

import (
        "engine-x/wal"
)

func (p *Pool) WritePageLog(
        pageID uint64,
        txid uint64,
        payload []byte,
        w *wal.WAL,
) error {

        idx,
                ok :=
                p.PageTable[pageID]

        if !ok {
                return ErrPageNotFound
        }

        lsn,
                err :=
                w.Append(
                        txid,
                        payload,
                )

        if err != nil {
                return err
        }

        p.Frames[idx].LSN = lsn

        return nil
}

func (p *Pool) FlushWithWAL(
        pageID uint64,
        w *wal.WAL,
) error {

        if err :=
                w.Sync(); err != nil {
                return err
        }

        return p.Flush(
                pageID,
        )
}
