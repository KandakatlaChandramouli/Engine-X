package core

import "testing"

func TestGet(
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
			[]byte("alpha"),
			[]byte("beta"),
		)

	if !ok {
		t.Fatal("insert failed")
	}

	key,
		val,
		ok :=
		Get(
			&p,
			0,
		)

	if !ok {
		t.Fatal("get failed")
	}

	if string(key) != "alpha" {
		t.Fatal("bad key")
	}

	if string(val) != "beta" {
		t.Fatal("bad value")
	}
}
