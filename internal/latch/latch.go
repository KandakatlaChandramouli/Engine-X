package latch

import "sync"

type Mode uint8

const (
        Read Mode = iota
        Write
)

type Latch struct {
        mu      sync.RWMutex
        readers int
        writer  bool
}
