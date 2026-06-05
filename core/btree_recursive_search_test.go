package core

import "testing"

func TestBTreeSearchMultiLevel(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var internal Page
	var leaf Page

	InitInternalPage(&root, 1)
	InitInternalPage(&internal, 2)
	InitPage(&leaf, 3)

	SetLeftChild(
		&root,
		2,
	)

	SetLeftChild(
		&internal,
		3,
	)

	pager.Add(&root)
	pager.Add(&internal)
	pager.Add(&leaf)

	found,
		ok :=
		BTreeSearch(
			pager,
			1,
			[]byte("abc"),
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 3 {
		t.Fatalf(
			"expected 3 got %d",
			PageID(found),
		)
	}
}
