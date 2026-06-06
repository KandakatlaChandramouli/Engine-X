package bufferpoolv2

type Stats struct {
        Hits   int
        Misses int
}

func (s *Stats) RecordHit() {
        s.Hits++
}

func (s *Stats) RecordMiss() {
        s.Misses++
}

func (s *Stats) HitRate() float64 {
        total := s.Hits + s.Misses

        if total == 0 {
                return 0
        }

        return float64(s.Hits) / float64(total)
}
