package mvcc

import "testing"

func TestVisibility(
        t *testing.T,
) {

        v := &Version{
                BeginTS: 10,
                EndTS:   20,
        }

        if !Visible(v,15) {
                t.Fatal()
        }

        if Visible(v,25) {
                t.Fatal()
        }
}
