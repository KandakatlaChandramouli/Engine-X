package core

import "testing"

func TestInternalEntryCount(
	t *testing.T,
) {

	var page Page

	InitInternalPage(
		&page,
		1,
	)

	SetLeftChild(
		&page,
		100,
	)

	if !InternalInsert(
		&page,
		[]byte("m"),
		200,
	) {
		t.Fatal()
	}

	if !InternalInsert(
		&page,
		[]byte("t"),
		300,
	) {
		t.Fatal()
	}

	count :=
		InternalEntryCount(
			&page,
		)

	if count != 2 {
		t.Fatalf(
			"expected 2 got %d",
			count,
		)
	}
}
