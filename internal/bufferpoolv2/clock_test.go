package bufferpoolv2

import "testing"

func TestClockAdd(t *testing.T) {
        c := NewClock()

        c.Add(1)
        c.Add(2)

        if c.Len() != 2 {
                t.Fatal()
        }
}

func TestClockVictim(t *testing.T) {
        c := NewClock()

        c.Add(1)
        c.Add(2)

        v, ok := c.Victim()

        if !ok || v != 1 {
                t.Fatal()
        }
}

func TestClockEmpty(t *testing.T) {
        c := NewClock()

        _, ok := c.Victim()

        if ok {
                t.Fatal()
        }
}
