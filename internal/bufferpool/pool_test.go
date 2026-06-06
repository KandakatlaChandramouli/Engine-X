package bufferpool

import "testing"

func TestCreatePool(
        t *testing.T,
) {

        p:=New(64)

        if len(p.Frames)!=64 {
                t.Fatal()
        }
}

func TestFetchPage(
        t *testing.T,
) {

        p:=New(4)

        f,err :=
                p.Fetch(100)

        if err != nil {
                t.Fatal(err)
        }

        if f.PageID != 100 {
                t.Fatal()
        }
}

func TestPinUnpin(
        t *testing.T,
) {

        p:=New(4)

        _,_ =
                p.Fetch(1)

        if err :=
                p.Unpin(1); err != nil {
                t.Fatal(err)
        }

        idx :=
                p.PageTable[1]

        if p.Frames[idx].Pin != 0 {
                t.Fatal()
        }
}

func TestDirtyPage(
        t *testing.T,
) {

        p:=New(4)

        _,_ =
                p.Fetch(1)

        if err :=
                p.MarkDirty(1); err != nil {
                t.Fatal(err)
        }

        idx :=
                p.PageTable[1]

        if !p.Frames[idx].Dirty {
                t.Fatal()
        }
}

func TestPoolFull(
        t *testing.T,
) {

        p:=New(2)

        _,_ = p.Fetch(1)
        _,_ = p.Fetch(2)

        _,err :=
                p.Fetch(3)

        if err != ErrPoolFull {
                t.Fatal()
        }
}
