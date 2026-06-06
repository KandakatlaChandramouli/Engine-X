package bufferpoolv2

import "testing"

func TestStatsHit(t *testing.T) {
        s := &Stats{}

        s.RecordHit()

        if s.Hits != 1 {
                t.Fatal()
        }
}

func TestStatsMiss(t *testing.T) {
        s := &Stats{}

        s.RecordMiss()

        if s.Misses != 1 {
                t.Fatal()
        }
}

func TestHitRate(t *testing.T) {
        s := &Stats{}

        s.RecordHit()
        s.RecordHit()
        s.RecordMiss()

        if s.HitRate() <= 0 {
                t.Fatal()
        }
}
