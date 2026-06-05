package core

import "testing"

func TestRecursiveRootCreation(
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
		[]byte("g"),
		300,
	)

	InternalInsert(
		&parent,
		[]byte("m"),
		400,
	)

	InternalInsert(
		&parent,
		[]byte("z"),
		500,
	)

	result,
		ok :=
		RecursivePropagate(
			&parent,
			&sibling,
		)

	if !ok {
		t.Fatal()
	}

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	if PageSlotCount(&parent) == 0 {
		t.Fatal()
	}

	if PageSlotCount(&sibling) == 0 {
		t.Fatal()
	}
}
