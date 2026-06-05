package core

import "testing"

func TestRecursivePropagation(
	t *testing.T,
) {

	var parent Page
	var sibling Page

	InitInternalPage(
		&parent,
		1,
	)

	SetLeftChild(
		&parent,
		100,
	)

	InternalInsert(
		&parent,
		[]byte("a"),
		200,
	)

	InternalInsert(
		&parent,
		[]byte("m"),
		300,
	)

	InternalInsert(
		&parent,
		[]byte("z"),
		400,
	)

	result,
		ok :=
		InternalSplit(
			&parent,
			&sibling,
		)

	if !ok {
		t.Fatal()
	}

	if string(result.SeparatorKey) != "m" {
		t.Fatal()
	}
}
