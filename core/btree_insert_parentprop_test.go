package core

import "testing"

func TestBTreeInsertParentPropagation(
	t *testing.T,
) {

	var parent Page
	var child Page

	InitInternalPage(
		&parent,
		1,
	)

	InitPage(
		&child,
		2,
	)

	SetParentPageID(
		&child,
		1,
	)

	result,
		ok :=
		ParentAwarePropagation(
			&parent,
			[]byte("m"),
			2,
		)

	if !ok {
		t.Fatal()
	}

	if string(result.SeparatorKey) != "m" {
		t.Fatal()
	}

	if result.RightPageID != 2 {
		t.Fatal()
	}

	if PageSlotCount(&parent) != 1 {
		t.Fatal()
	}
}
