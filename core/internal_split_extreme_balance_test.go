package core

import "testing"

func TestInternalSplitExtremeBalance(
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
		999,
	)

	inserted := 0

	for i := 0; i < 256; i++ {

		if !InternalInsert(
			&left,
			[]byte{
				byte(i),
				byte(i >> 1),
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
		int(
			InternalEntryCount(
				&left,
			),
		)

	rightCount :=
		int(
			InternalEntryCount(
				&right,
			),
		)

	diff := leftCount - rightCount

	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		t.Fatalf(
			"unbalanced split: left=%d right=%d",
			leftCount,
			rightCount,
		)
	}

	if leftCount+rightCount+1 != inserted {
		t.Fatal()
	}
}
