package core

import "testing"

func TestBTreeGetMissingKey(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var leaf Page

	InitInternalPage(&root, 1)
	InitPage(&leaf, 2)

	SetLeftChild(
		&root,
		2,
	)

	if !LeafInsert(
		&leaf,
		[]byte("hello"),
		[]byte("world"),
	) {
		t.Fatal()
	}

	pager.Add(&root)
	pager.Add(&leaf)

	_,
		ok :=
		BTreeGet(
			pager,
			1,
			[]byte("missing"),
		)

	if ok {
		t.Fatal()
	}
}
