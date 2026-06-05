package core

import "testing"

func TestBTreeSearchFindsLeafRecord(
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

	key,
		value,
		ok :=
		Get(
			found,
			0,
		)

	if !ok {
		t.Fatal()
	}

	if string(key) != "hello" {
		t.Fatal()
	}

	if string(value) != "world" {
		t.Fatal()
	}
}
