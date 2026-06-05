package core

import "testing"

func TestInternalSplitProducesStableStructure(
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

	for i := 0; i < 64; i++ {

		if !InternalInsert(
			&left,
			[]byte{
				byte(i),
			},
			uint64(i+1000),
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

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	leftCount :=
		InternalEntryCount(
			&left,
		)

	rightCount :=
		InternalEntryCount(
			&right,
		)

	if leftCount == 0 {
		t.Fatal()
	}

	if rightCount == 0 {
		t.Fatal()
	}

	if LeftChild(
		&left,
	) == 0 {
		t.Fatal()
	}

	if LeftChild(
		&right,
	) == 0 {
		t.Fatal()
	}
}
