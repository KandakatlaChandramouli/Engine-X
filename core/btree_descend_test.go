package core

import "testing"

func TestBTreeSearchDescendsInternal(
	t *testing.T,
) {

	pager :=
		NewPager()

	var root Page
	var left Page
	var right Page

	InitInternalPage(
		&root,
		1,
	)

	InitPage(
		&left,
		2,
	)

	InitPage(
		&right,
		3,
	)

	SetLeftChild(
		&root,
		2,
	)

	if !InternalInsert(
		&root,
		[]byte("m"),
		3,
	) {
		t.Fatal()
	}

	pager.Add(&root)
	pager.Add(&left)
	pager.Add(&right)

	found,
		ok :=
		BTreeSearch(
			pager,
			1,
			[]byte("apple"),
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 2 {
		t.Fatalf(
			"expected 2 got %d",
			PageID(found),
		)
	}
}
