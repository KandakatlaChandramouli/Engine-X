package core

import "testing"

func TestAllocatePageID(
	t *testing.T,
) {

	pager := NewPager()

	var p1 Page
	var p2 Page

	InitPage(
		&p1,
		1,
	)

	InitPage(
		&p2,
		5,
	)

	pager.Add(&p1)
	pager.Add(&p2)

	id :=
		pager.AllocatePageID()

	if id != 6 {
		t.Fatalf(
			"expected 6 got %d",
			id,
		)
	}
}
