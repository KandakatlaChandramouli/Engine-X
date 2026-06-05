package core

import "testing"

func TestAllocatePageIDMonotonic(
	t *testing.T,
) {

	pager := NewPager()

	var p1 Page
	var p2 Page
	var p3 Page

	InitPage(&p1, 1)
	InitPage(&p2, 5)
	InitPage(&p3, 9)

	pager.Add(&p1)
	pager.Add(&p2)
	pager.Add(&p3)

	id1 :=
		pager.AllocatePageID()

	var p4 Page

	InitPage(
		&p4,
		id1,
	)

	pager.Add(&p4)

	id2 :=
		pager.AllocatePageID()

	if id2 <= id1 {
		t.Fatalf(
			"expected id2 > id1, got %d <= %d",
			id2,
			id1,
		)
	}
}
