package bufferpool

import "errors"

var (
        ErrPageNotFound = errors.New(
                "page not found",
        )

        ErrPoolFull = errors.New(
                "buffer pool full",
        )
)

type Pool struct {
        Frames    []Frame
        PageTable map[uint64]int
}

func New(
        size int,
) *Pool {

        return &Pool{
                Frames: make(
                        []Frame,
                        size,
                ),
                PageTable: make(
                        map[uint64]int,
                ),
        }
}

func (p *Pool) Fetch(
        pageID uint64,
) (*Frame,error) {

        idx,ok :=
                p.PageTable[pageID]

        if ok {
                p.Frames[idx].Pin++
                return &p.Frames[idx],nil
        }

        for i:=0;i<len(p.Frames);i++ {

                if p.Frames[i].Used {
                        continue
                }

                p.Frames[i].Used=true
                p.Frames[i].PageID=pageID
                p.Frames[i].Pin=1

                p.PageTable[pageID]=i

                return &p.Frames[i],nil
        }

        return nil,
                ErrPoolFull
}

func (p *Pool) Unpin(
        pageID uint64,
) error {

        idx,ok :=
                p.PageTable[pageID]

        if !ok {
                return ErrPageNotFound
        }

        if p.Frames[idx].Pin > 0 {
                p.Frames[idx].Pin--
        }

        return nil
}

func (p *Pool) MarkDirty(
        pageID uint64,
) error {

        idx,ok :=
                p.PageTable[pageID]

        if !ok {
                return ErrPageNotFound
        }

        p.Frames[idx].Dirty=true

        return nil
}
