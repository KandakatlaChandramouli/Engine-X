package core

import "testing"

func TestInternalSplitResult(
	t *testing.T,
) {

	r :=
		InternalSplitResult{
			SeparatorKey: []byte("m"),
		}

	if string(r.SeparatorKey) != "m" {
		t.Fatal()
	}
}
