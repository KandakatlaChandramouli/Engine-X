package core

import "testing"

func TestHandleParentOverflow(
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

	for i := 0; i < 200; i++ {

		ok :=
			InternalInsert(
				&parent,
				[]byte(string(rune('a'+(i%26)))),
				uint64(i+1000),
			)

		if !ok {
			break
		}
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
}
