package core

import "testing"

func TestBTreeSearchRoutesBySeparator(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var left Page
	var right Page

	InitInternalPage(&root, 1)

	InitPage(&left, 2)
	InitPage(&right, 3)

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

	page,
		ok :=
		BTreeSearch(
			pager,
			1,
			[]byte("apple"),
		)

	if !ok {
		t.Fatal()
	}

	if PageID(page) != 2 {
		t.Fatalf(
			"expected left page got %d",
			PageID(page),
		)
	}

	page,
		ok =
		BTreeSearch(
			pager,
			1,
			[]byte("zebra"),
		)

	if !ok {
		t.Fatal()
	}

	if PageID(page) != 3 {
		t.Fatalf(
			"expected right page got %d",
			PageID(page),
		)
	}
}
