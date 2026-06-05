package core

import "testing"

func TestInternalSplit(
	t *testing.T,
) {

	var left Page
	var right Page

	InitInternalPage(
		&left,
		1,
	)

	SetLeftChild(
		&left,
		100,
	)

	if !InternalInsert(
		&left,
		[]byte("a"),
		200,
	) {
		t.Fatal()
	}

	if !InternalInsert(
		&left,
		[]byte("m"),
		300,
	) {
		t.Fatal()
	}

	if !InternalInsert(
		&left,
		[]byte("z"),
		400,
	) {
		t.Fatal()
	}

	result,
		ok :=
		InternalSplit(
			&left,
			&right,
		)

	if !ok {
		t.Fatal()
	}

	if string(result.SeparatorKey) != "m" {
		t.Fatalf(
			"expected m got %s",
			string(result.SeparatorKey),
		)
	}

	if PageSlotCount(&left) != 1 {
		t.Fatal()
	}

	if PageSlotCount(&right) != 1 {
		t.Fatal()
	}
}
