package core

import "testing"

func TestInternalSplitChildCounts(
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
		500,
	)

	inserted := 0

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

	if len(
		result.SeparatorKey,
	) == 0 {
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

	if int(leftCount+rightCount+1) != inserted {
		t.Fatalf(
			"expected %d got %d",
			inserted,
			leftCount+rightCount+1,
		)
	}
}
