package core

import "testing"

func TestInternalSearchExactSeparator(
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

	if !ok {
		t.Fatal()
	}

	if child != 200 {
		t.Fatalf(
			"expected 200 got %d",
			child,
		)
	}
}
