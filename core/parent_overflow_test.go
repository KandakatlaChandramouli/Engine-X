package core

import "testing"

func TestParentOverflow(
	t *testing.T,
) {

	var parent Page

	InitInternalPage(
		&parent,
		1,
	)

	SetLeftChild(
		&parent,
		100,
	)

	overflow := false

	for i := 0; i < 1000; i++ {

		if ParentOverflow(
			&parent,
			[]byte(string(rune('a'+(i%26)))),
			uint64(i+200),
		) {
			overflow = true
			break
		}
	}

	if !overflow {
		t.Fatal()
	}
}
