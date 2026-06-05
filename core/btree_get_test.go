package core

import "testing"

func TestBTreeGet(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var leaf Page

	InitInternalPage(
		&root,
		1,
	)

	InitPage(
		&leaf,
		2,
	)

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

	value,
		ok :=
		BTreeGet(
			pager,
			1,
			[]byte("hello"),
		)

	if !ok {
		t.Fatal()
	}

	if string(value) != "world" {
		t.Fatal()
	}
}
