package core

import "testing"

func TestMissingKeyLookup(
	t *testing.T,
) {

	var page Page

	InitPage(
		&page,
		1,
	)

	if !LeafInsert(
		&page,
		[]byte("alpha"),
		[]byte("value"),
	) {
		t.Fatal()
	}

	pos :=
		LeafSearch(
			&page,
			[]byte("missing"),
		)

	if pos >= 0 {
		t.Fatal()
	}
}
