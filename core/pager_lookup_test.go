package core

import "testing"

func TestPagerLookupMultiplePages(
	t *testing.T,
) {

	pager :=
		NewPager()

	var p1 Page
	var p2 Page
	var p3 Page

	InitPage(&p1, 1)
	InitPage(&p2, 2)
	InitPage(&p3, 3)

	pager.Add(&p1)
	pager.Add(&p2)
	pager.Add(&p3)

	_, ok := pager.Get(1)

	if !ok {
		t.Fatal()
	}

	_, ok = pager.Get(2)

	if !ok {
		t.Fatal()
	}

	_, ok = pager.Get(3)

	if !ok {
		t.Fatal()
	}

	_, ok = pager.Get(99)

	if ok {
		t.Fatal()
	}
}
