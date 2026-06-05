package core

import "testing"

func TestBTreeInsertRealSplit(
	t *testing.T,
) {
	var left Page
	var right Page

	InitPage(&left, 1)
	InitPage(&right, 2)

	inserted := 0

	for {
		ok := LeafInsert(
			&left,
			[]byte(string(rune('a'+(inserted%26)))),
			[]byte("value"),
		)

		if !ok {
			break
		}

		inserted++
	}

	result, ok :=
		LeafSplit(
			&left,
			&right,
		)

	if !ok {
		t.Fatal()
	}

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	if PageSlotCount(&left) == 0 {
		t.Fatal()
	}

	if PageSlotCount(&right) == 0 {
		t.Fatal()
	}
}
