package core

import "testing"

func TestBTreeInsert(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	result,
		ok :=
		BTreeInsert(
			&p,
			[]byte("a"),
			[]byte("1"),
		)

	if !ok {
		t.Fatal()
	}

	if result.Split {
		t.Fatal()
	}

	k,
		v,
		ok :=
		Get(
			&p,
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
