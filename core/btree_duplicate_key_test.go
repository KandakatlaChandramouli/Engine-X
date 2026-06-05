package core

import "testing"

func TestDuplicateKeyRejected(
	t *testing.T,
) {

	var page Page

	InitPage(
		&page,
		1,
	)

	if !LeafInsert(
		&page,
		[]byte("alpha"),
		[]byte("v1"),
	) {
		t.Fatal()
	}

	if LeafInsert(
		&page,
		[]byte("alpha"),
		[]byte("v2"),
	) {
		t.Fatal()
	}

	value,
		ok :=
		Get(
			&page,
			[]byte("alpha"),
		)

	if !ok {
		t.Fatal()
	}

	if string(value) != "v1" {
		t.Fatal()
	}
}
