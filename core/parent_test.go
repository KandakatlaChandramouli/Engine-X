package core

import "testing"

func TestParentPageID(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	SetParentPageID(
		&p,
		99,
	)

	if ParentPageID(&p) != 99 {
		t.Fatal()
	}
}
