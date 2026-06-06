package bufferpoolv2

type Prefetcher struct {
        queue []int
}

func NewPrefetcher() *Prefetcher {
        return &Prefetcher{}
}

func (p *Prefetcher) Add(pageID int) {
        p.queue = append(p.queue, pageID)
}

func (p *Prefetcher) Next() (int, bool) {
        if len(p.queue) == 0 {
                return 0, false
        }

        v := p.queue[0]
        p.queue = p.queue[1:]

        return v, true
}

func (p *Prefetcher) Len() int {
        return len(p.queue)
}
