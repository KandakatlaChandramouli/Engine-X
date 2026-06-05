package core

import "testing"

func TestLeafInsertOrdered(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	SetPageType(
		&p,
		PageTypeLeaf,
	)

	if !LeafInsert(
		&p,
		[]byte("zebra"),
		[]byte("1"),
	) {
		t.Fatal()
	}

	if !LeafInsert(
		&p,
		[]byte("apple"),
		[]byte("2"),
	) {
		t.Fatal()
	}

	if !LeafInsert(
		&p,
		[]byte("banana"),
		[]byte("3"),
	) {
		t.Fatal()
	}

	k0, _, _ := Get(&p, 0)
	k1, _, _ := Get(&p, 1)
	k2, _, _ := Get(&p, 2)

	if string(k0) != "apple" {
		t.Fatal(string(k0))
	}

	if string(k1) != "banana" {
		t.Fatal(string(k1))
	}

	if string(k2) != "zebra" {
		t.Fatal(string(k2))
	}
}

func TestLeafGetOrdered(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	SetPageType(
		&p,
		PageTypeLeaf,
	)

	LeafInsert(
		&p,
		[]byte("zebra"),
		[]byte("1"),
	)

	LeafInsert(
		&p,
		[]byte("apple"),
		[]byte("2"),
	)

	LeafInsert(
		&p,
		[]byte("banana"),
		[]byte("3"),
	)

	v,
		ok :=
		LeafGet(
			&p,
			[]byte("banana"),
		)

	if !ok {
		t.Fatal()
	}

	if string(v) != "3" {
		t.Fatal(string(v))
	}
}
