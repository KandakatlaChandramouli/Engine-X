package core

import (
	"bytes"
	"testing"
)

func TestInternalSplitRoundTrip(
	t *testing.T,
) {

	var original Page
	var left Page
	var right Page

	InitInternalPage(
		&original,
		1,
	)

	InitInternalPage(
		&left,
		2,
	)

	InitInternalPage(
		&right,
		3,
	)

	SetLeftChild(
		&original,
		100,
	)

	inserted := 0

	for i := 0; i < 256; i++ {

		if !InternalInsert(
			&original,
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
			&original,
			&right,
		)

	if !ok {
		t.Fatal()
	}

	left = original

	total :=
		InternalEntryCount(
			&left,
		) +
			InternalEntryCount(
				&right,
			)

	if total+1 < uint16(inserted) {
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
}
