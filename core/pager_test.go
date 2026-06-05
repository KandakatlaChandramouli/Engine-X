package core

import "testing"

func TestPager(
	t *testing.T,
) {

	pager :=
		NewPager()

	var page Page

	InitPage(
		&page,
		99,
	)

	pager.Add(
		&page,
	)

	found,
		ok :=
		pager.Get(
			99,
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 99 {
		t.Fatal()
	}
}
