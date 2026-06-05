package core

import "testing"

func TestInternalSearchBoundaryKeys(
	t *testing.T,
) {

	var page Page

	InitInternalPage(
		&page,
		1,
	)

	SetLeftChild(
		&page,
		100,
	)

	InternalInsert(
		&page,
		[]byte("m"),
		200,
	)

	InternalInsert(
		&page,
		[]byte("t"),
		300,
	)

	child,
		ok :=
		InternalSearch(
			&page,
			[]byte("m"),
		)

	if !ok || child != 200 {
		t.Fatal()
	}

	child,
		ok =
		InternalSearch(
			&page,
			[]byte("t"),
		)

	if !ok || child != 300 {
		t.Fatal()
	}
}
