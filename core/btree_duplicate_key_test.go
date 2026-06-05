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

	pos :=
		LeafSearch(
			&page,
			[]byte("alpha"),
		)

	if pos < 0 {
		t.Fatal()
	}

	key,
		value,
		ok :=
		Get(
			&page,
			uint16(pos),
		)

	if !ok {
		t.Fatal()
	}

	if string(key) != "alpha" {
		t.Fatal()
	}

	if string(value) != "v1" {
		t.Fatal()
	}
}
