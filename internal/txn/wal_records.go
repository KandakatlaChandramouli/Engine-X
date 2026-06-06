package txn

import (
        "engine-x/wal"
)

var (
        BeginRecord  = []byte("BEGIN")
        CommitRecord = []byte("COMMIT")
        AbortRecord  = []byte("ABORT")
)

func (m *Manager) Commit(
        tx *Transaction,
        w *wal.WAL,
) error {

        lsn,
                err :=
                w.Append(
                        tx.ID,
                        CommitRecord,
                )

        if err != nil {
                return err
        }

        tx.State = Committed
        tx.EndLSN = lsn

        m.mu.Lock()
        delete(
                m.active,
                tx.ID,
        )
        m.mu.Unlock()

        return nil
}

func (m *Manager) Abort(
        tx *Transaction,
        w *wal.WAL,
) error {

        lsn,
                err :=
                w.Append(
                        tx.ID,
                        AbortRecord,
                )

        if err != nil {
                return err
        }

        tx.State = Aborted
        tx.EndLSN = lsn

        m.mu.Lock()
        delete(
                m.active,
                tx.ID,
        )
        m.mu.Unlock()

        return nil
}
