package core

import "testing"

func TestInsertRecursive(
	t *testing.T,
) {

	pager := NewPager()

	var leaf Page

	InitPage(
		&leaf,
		1,
	)

	pager.Add(
		&leaf,
	)

	result,
		ok :=
		InsertRecursive(
			pager,
			1,
			[]byte("hello"),
			[]byte("world"),
		)

	if !ok {
		t.Fatal()
	}

	if result.Split {
		t.Fatal()
	}

	value,
		ok :=
		LeafGet(
			&leaf,
			[]byte("hello"),
		)

	if !ok {
		t.Fatal()
	}

	if string(value) != "world" {
		t.Fatal()
	}
}
