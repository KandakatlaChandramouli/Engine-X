package bufferpoolv2

import "testing"

func TestLRUPush(
        t *testing.T,
) {
        l := NewLRU()

        l.Touch(1)
        l.Touch(2)

        if l.Len() != 2 {
                t.Fatal()
        }
}

func TestLRUEvict(
        t *testing.T,
) {
        l := NewLRU()

        l.Touch(1)
        l.Touch(2)

        v, ok := l.Victim()

        if !ok || v != 1 {
                t.Fatal()
        }
}
