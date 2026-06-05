package core

import "testing"

func TestBTreeInsert(
	t *testing.T,
) {
	pager := NewPager()

	var root Page

	InitPage(
		&root,
		1,
	)

	pager.Add(
		&root,
	)

	ok :=
		BTreeInsert(
			pager,
			&root,
			[]byte("a"),
			[]byte("1"),
		)

	if !ok {
		t.Fatal()
	}

	k,
		v,
		ok :=
		Get(
			&root,
			0,
		)

	if !ok {
		t.Fatal()
	}

	if string(k) != "a" {
		t.Fatal()
	}

	if string(v) != "1" {
		t.Fatal()
	}
}
