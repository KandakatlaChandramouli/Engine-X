package core

import "testing"

func TestInternalSearchLeftChild(
	t *testing.T,
) {

	var p Page

	InitInternalPage(
		&p,
		1,
	)

	SetLeftChild(
		&p,
		100,
	)

	if !InternalInsert(
		&p,
		[]byte("m"),
		200,
	) {
		t.Fatal()
	}

	child,
		ok :=
		InternalSearch(
			&p,
			[]byte("apple"),
		)

	if !ok {
		t.Fatal()
	}

	if child != 100 {
		t.Fatalf(
			"expected 100 got %d",
			child,
		)
	}
}
