package bufferpoolv2

import "sync"

type Prefetcher struct {
	mu    sync.Mutex
	queue []int
}

func NewPrefetcher() *Prefetcher {
	return &Prefetcher{}
}

func (p *Prefetcher) Add(page int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.queue = append(p.queue, page)
}

func (p *Prefetcher) Next() (int, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.queue) == 0 {
		return 0, false
	}

	v := p.queue[0]
	p.queue = p.queue[1:]

	return v, true
}

func (p *Prefetcher) Len() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.queue)
}
