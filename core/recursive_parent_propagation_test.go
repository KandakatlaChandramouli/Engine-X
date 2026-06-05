package core

import "testing"

func TestRecursiveParentPropagation(
	t *testing.T,
) {

	var parent Page
	var sibling Page

	InitInternalPage(
		&parent,
		1,
	)

	InitInternalPage(
		&sibling,
		2,
	)

	SetLeftChild(
		&parent,
		100,
	)

	child := uint64(200)

	for {

		if !InternalInsert(
			&parent,
			[]byte(string(rune('a'+(child%26)))),
			child,
		) {
			break
		}

		child++
	}

	result,
		ok :=
		HandleParentOverflow(
			&parent,
			&sibling,
		)

	if !ok {
		t.Fatal()
	}

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	if InternalEntryCount(
		&sibling,
	) == 0 {
		t.Fatal()
	}
}
