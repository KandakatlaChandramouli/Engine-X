package bufferpool

import "testing"

func TestCreatePool(
        t *testing.T,
) {

        p := New(64)

        if len(
                p.Frames,
        ) != 64 {

                t.Fatal()
        }
}
