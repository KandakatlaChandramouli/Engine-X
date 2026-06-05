package core

import "testing"

func TestPagerRootLookup(
	t *testing.T,
) {

	pager :=
		NewPager()

	var root Page

	InitInternalPage(
		&root,
		1,
	)

	pager.Add(
		&root,
	)

	found,
		ok :=
		PagerRootLookup(
			pager,
			1,
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 1 {
		t.Fatal()
	}
}
