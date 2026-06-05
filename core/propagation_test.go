package core

import "testing"

func TestPropagationResult(
	t *testing.T,
) {

	r :=
		PropagationResult{
			SeparatorKey: []byte("m"),
			RightPageID:  99,
		}

	if string(r.SeparatorKey) != "m" {
		t.Fatal()
	}

	if r.RightPageID != 99 {
		t.Fatal()
	}
}
