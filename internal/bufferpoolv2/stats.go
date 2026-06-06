package bufferpoolv2

import "sync/atomic"

type Stats struct {
	Hits   int64
	Misses int64
}

func (s *Stats) RecordHit() {
	atomic.AddInt64(&s.Hits, 1)
}

func (s *Stats) RecordMiss() {
	atomic.AddInt64(&s.Misses, 1)
}

func (s *Stats) HitRate() float64 {
	hits := atomic.LoadInt64(&s.Hits)
	misses := atomic.LoadInt64(&s.Misses)

	total := hits + misses

	if total == 0 {
		return 0
	}

	return float64(hits) / float64(total)
}
