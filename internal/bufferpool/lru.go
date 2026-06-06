package bufferpool

import "errors"

var ErrNoVictim = errors.New(
        "no victim available",
)

func (p *Pool) Victim() (
        int,
        error,
) {

        oldest := -1

        for i := 0; i < len(p.Frames); i++ {

                if !p.Frames[i].Used {
                        continue
                }

                if p.Frames[i].Pin != 0 {
                        continue
                }

                oldest = i
                break
        }

        if oldest == -1 {
                return -1,
                        ErrNoVictim
        }

        return oldest,
                nil
}

func (p *Pool) Evict() error {

        victim,
                err :=
                p.Victim()

        if err != nil {
                return err
        }

        delete(
                p.PageTable,
                p.Frames[victim].PageID,
        )

        p.Frames[victim] = Frame{}

        return nil
}
