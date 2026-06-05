package core

import "testing"

func TestRecursiveSplit(
	t *testing.T,
) {

	var page Page
	var sibling Page

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
		[]byte("a"),
		200,
	)

	InternalInsert(
		&page,
		[]byte("g"),
		300,
	)

	InternalInsert(
		&page,
		[]byte("m"),
		400,
	)

	InternalInsert(
		&page,
		[]byte("z"),
		500,
	)

	result,
		ok :=
		RecursiveSplit(
			&page,
			&sibling,
		)

	if !ok {
		t.Fatal()
	}

	if string(result.SeparatorKey) != "m" {
		t.Fatal()
	}
}
