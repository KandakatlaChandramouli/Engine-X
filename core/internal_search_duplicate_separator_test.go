package core

import "testing"

func TestInternalSearchDuplicateSeparatorBehavior(
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

	if InternalInsert(
		&page,
		[]byte("m"),
		300,
	) {
		t.Fatal()
	}

	count :=
		InternalEntryCount(
			&page,
		)

	if count != 1 {
		t.Fatalf(
			"expected 1 got %d",
			count,
		)
	}
}
