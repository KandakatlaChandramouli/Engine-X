package core

import "testing"

func TestParentAwarePropagation(
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

	if ParentPageID(&child) != 99 {
		t.Fatal()
	}
}
