package core

import "testing"

func TestInsertIntoParent(
	t *testing.T,
) {

	var parent Page

	InitInternalPage(
		&parent,
		1,
	)

	SetLeftChild(
		&parent,
		100,
	)

	if !InternalInsert(
		&parent,
		[]byte("m"),
		200,
	) {
		t.Fatal()
	}

	if !InternalInsert(
		&parent,
		[]byte("z"),
		300,
	) {
		t.Fatal()
	}

	if !InsertIntoParent(
		&parent,
		[]byte("g"),
		400,
	) {
		t.Fatal()
	}

	if PageSlotCount(&parent) != 3 {
		t.Fatal()
	}

	e0, _ := ReadInternalEntry(&parent, 0)
	e1, _ := ReadInternalEntry(&parent, 1)
	e2, _ := ReadInternalEntry(&parent, 2)

	if string(e0.Key) != "g" {
		t.Fatal(string(e0.Key))
	}

	if string(e1.Key) != "m" {
		t.Fatal(string(e1.Key))
	}

	if string(e2.Key) != "z" {
		t.Fatal(string(e2.Key))
	}
}
