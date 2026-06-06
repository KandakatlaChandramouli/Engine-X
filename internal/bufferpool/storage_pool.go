package bufferpool

import (
        "errors"

        "engine-x/core"
)

var ErrNilDB = errors.New(
        "nil db",
)

func (p *Pool) FetchFromStorage(
        db *core.DB,
        pageID uint64,
) (*Frame,error) {

        if db == nil {
                return nil,
                        ErrNilDB
        }

        frame,
                err :=
                p.Fetch(
                        pageID,
                )

        if err != nil {
                return nil,
                        err
        }

        frame.DB = db

        if frame.Data == nil {

                frame.Data =
                        db.Page(
                                pageID,
                        )
        }

        return frame,
                nil
}

func (p *Pool) Flush(
        pageID uint64,
) error {

        idx,
                ok :=
                p.PageTable[pageID]

        if !ok {
                return ErrPageNotFound
        }

        frame :=
                &p.Frames[idx]

        if frame.DB == nil {
                return ErrNilDB
        }

        if !frame.Dirty {
                return nil
        }

        copy(
                frame.DB.Page(
                        pageID,
                ),
                frame.Data,
        )

        frame.Dirty = false

        return frame.DB.Sync()
}

func (p *Pool) FlushAll() error {

        for pageID :=
                range p.PageTable {

                if err :=
                        p.Flush(
                                pageID,
                        ); err != nil {

                        return err
                }
        }

        return nil
}
