package core

import (
	"bytes"
	"testing"
)

func TestInternalSplitSeparatorOrdering(
	t *testing.T,
) {

	var left Page
	var right Page

	InitInternalPage(
		&left,
		1,
	)

	InitInternalPage(
		&right,
		2,
	)

	SetLeftChild(
		&left,
		100,
	)

	keys := [][]byte{
		[]byte("a"),
		[]byte("c"),
		[]byte("e"),
		[]byte("g"),
		[]byte("i"),
		[]byte("k"),
		[]byte("m"),
		[]byte("o"),
	}

	for i, key := range keys {

		if !InternalInsert(
			&left,
			key,
			uint64(i+1000),
		) {
			t.Fatal()
		}
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

	leftLast,
		ok :=
		ReadInternalEntry(
			&left,
			InternalEntryCount(&left)-1,
		)

	if !ok {
		t.Fatal()
	}

	if bytes.Compare(
		leftLast.Key,
		result.SeparatorKey,
	) >= 0 {
		t.Fatal()
	}

	rightFirst,
		ok :=
		ReadInternalEntry(
			&right,
			0,
		)

	if !ok {
		t.Fatal()
	}

	if bytes.Compare(
		rightFirst.Key,
		result.SeparatorKey,
	) <= 0 {
		t.Fatal()
	}
}
