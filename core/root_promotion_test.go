package core

import "testing"

func TestPromoteRoot(
	t *testing.T,
) {

	var root Page

	InitPage(
		&root,
		99,
	)

	ok :=
		PromoteRoot(
			&root,
			1,
			2,
			[]byte("m"),
		)

	if !ok {
		t.Fatal()
	}

	if PageType(&root) != PageTypeInternal {
		t.Fatal()
	}

	if LeftChild(&root) != 1 {
		t.Fatal()
	}

	child,
		ok :=
		InternalSearch(
			&root,
			[]byte("apple"),
		)

	if !ok {
		t.Fatal()
	}

	if child != 1 {
		t.Fatalf(
			"expected 1 got %d",
			child,
		)
	}

	child,
		ok =
		InternalSearch(
			&root,
			[]byte("zebra"),
		)

	if !ok {
		t.Fatal()
	}

	if child != 2 {
		t.Fatalf(
			"expected 2 got %d",
			child,
		)
	}
}
