package core

import (
	"bytes"
	"testing"
)

func TestInternalSplitSearchability(
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

	inserted := 0

	for i := 0; i < 256; i++ {

		if !InternalInsert(
			&left,
			[]byte{
				byte(i),
				byte(i + 1),
			},
			uint64(i+1000),
		) {
			break
		}

		inserted++
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

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	leftCount :=
		InternalEntryCount(
			&left,
		)

	for i := uint16(0); i < leftCount; i++ {

		entry,
			ok :=
			ReadInternalEntry(
				&left,
				i,
			)

		if !ok {
			t.Fatal()
		}

		if bytes.Compare(
			entry.Key,
			result.SeparatorKey,
		) >= 0 {
			t.Fatal()
		}
	}

	rightCount :=
		InternalEntryCount(
			&right,
		)

	for i := uint16(0); i < rightCount; i++ {

		entry,
			ok :=
			ReadInternalEntry(
				&right,
				i,
			)

		if !ok {
			t.Fatal()
		}

		if bytes.Compare(
			entry.Key,
			result.SeparatorKey,
		) < 0 {
			t.Fatal()
		}
	}

	_ = inserted
}
