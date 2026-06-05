package core

import "testing"

func TestPagerParentLookup(
	t *testing.T,
) {

	pager :=
		NewPager()

	var parent Page
	var child Page

	InitInternalPage(
		&parent,
		99,
	)

	InitPage(
		&child,
		100,
	)

	SetParentPageID(
		&child,
		99,
	)

	pager.Add(
		&parent,
	)

	pager.Add(
		&child,
	)

	found,
		ok :=
		PagerParentLookup(
			pager,
			&child,
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 99 {
		t.Fatal()
	}
}
