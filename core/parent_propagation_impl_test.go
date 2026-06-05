package core

import "testing"

func TestParentAwarePropagationImpl(
	t *testing.T,
) {

	var child Page

	InitPage(
		&child,
		10,
	)

	SetParentPageID(
		&child,
		99,
	)

	parent,
		ok :=
		ParentAwarePropagation(
			&child,
		)

	if !ok {
		t.Fatal()
	}

	if parent != 99 {
		t.Fatal()
	}
}
