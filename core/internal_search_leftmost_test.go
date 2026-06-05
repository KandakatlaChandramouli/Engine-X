package core

import "testing"

func TestInternalSearchLeftmostRouting(
	t *testing.T,
) {

	var page Page

	InitInternalPage(
		&page,
		1,
	)

	SetLeftChild(
		&page,
		111,
	)

	InternalInsert(
		&page,
		[]byte("m"),
		222,
	)

	child,
		ok :=
		InternalSearch(
			&page,
			[]byte("a"),
		)

	if !ok {
		t.Fatal()
	}

	if child != 111 {
		t.Fatal()
	}
}
