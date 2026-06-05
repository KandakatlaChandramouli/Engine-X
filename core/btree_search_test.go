package core

import "testing"

func TestBTreeSearchRoot(
	t *testing.T,
) {

	pager :=
		NewPager()

	var root Page

	InitPage(
		&root,
		1,
	)

	pager.Add(
		&root,
	)

	found,
		ok :=
		BTreeSearch(
			pager,
			1,
			[]byte("hello"),
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 1 {
		t.Fatal()
	}
}
