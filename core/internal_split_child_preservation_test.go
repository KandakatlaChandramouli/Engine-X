package core

import "testing"

func TestInternalSplitPreservesChildren(
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
		111,
	)

	for i := 0; i < 32; i++ {

		if !InternalInsert(
			&left,
			[]byte{
				byte('a' + i),
			},
			uint64(1000+i),
		) {
			break
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

	if len(
		result.SeparatorKey,
	) == 0 {
		t.Fatal()
	}

	if LeftChild(
		&left,
	) != 111 {
		t.Fatal()
	}

	if LeftChild(
		&right,
	) == 0 {
		t.Fatal()
	}

	if InternalEntryCount(
		&left,
	) == 0 {
		t.Fatal()
	}

	if InternalEntryCount(
		&right,
	) == 0 {
		t.Fatal()
	}
}
