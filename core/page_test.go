package core

import "testing"

func TestPageInsert(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	ok :=
		Insert(
			&p,
			[]byte("hello"),
			[]byte("world"),
		)

	if !ok {
		t.Fatal("insert failed")
	}

	if PageID(&p) != 1 {
		t.Fatal("bad page id")
	}
}
