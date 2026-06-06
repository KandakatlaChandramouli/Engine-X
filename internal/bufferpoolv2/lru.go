package bufferpoolv2

type LRU struct {
        pages []int
}

func NewLRU() *LRU {
        return &LRU{}
}

func (l *LRU) Touch(id int) {
        for i, v := range l.pages {
                if v == id {
                        l.pages = append(
                                l.pages[:i],
                                l.pages[i+1:]...,
                        )
                        break
                }
        }

        l.pages = append(l.pages, id)
}

func (l *LRU) Victim() (int, bool) {
        if len(l.pages) == 0 {
                return 0, false
        }

        v := l.pages[0]

        l.pages = l.pages[1:]

        return v, true
}

func (l *LRU) Len() int {
        return len(l.pages)
}
