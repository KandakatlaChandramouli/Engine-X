package core

import "testing"

func TestFreelist(
	t *testing.T,
) {

	var p Page

	InitFreelist(&p)

	if FreelistCount(&p) != 0 {
		t.Fatal("bad count")
	}

	ok :=
		FreelistPush(
			&p,
			100,
		)

	if !ok {
		t.Fatal("push failed")
	}

	ok =
		FreelistPush(
			&p,
			200,
		)

	if !ok {
		t.Fatal("push failed")
	}

	if FreelistCount(&p) != 2 {
		t.Fatal("bad count")
	}

	id,
		ok :=
		FreelistPop(
			&p,
		)

	if !ok {
		t.Fatal("pop failed")
	}

	if id != 200 {
		t.Fatal("bad page id")
	}

	id,
		ok =
		FreelistPop(
			&p,
		)

	if !ok {
		t.Fatal("pop failed")
	}

	if id != 100 {
		t.Fatal("bad page id")
	}
}
