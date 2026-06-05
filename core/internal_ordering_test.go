package core

import "testing"

func TestInternalOrdering(
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

	keys := [][]byte{
		[]byte("z"),
		[]byte("m"),
		[]byte("a"),
		[]byte("t"),
		[]byte("b"),
	}

	child := uint64(200)

	for _, key := range keys {

		if !InternalInsert(
			&page,
			key,
			child,
		) {
			t.Fatal()
		}

		child++
	}

	previous := ""

	count :=
		InternalEntryCount(
			&page,
		)

	for i := uint16(0); i < count; i++ {

		entry,
			ok :=
			ReadInternalEntry(
				&page,
				i,
			)

		if !ok {
			t.Fatal()
		}

		if previous != "" &&
			string(entry.Key) < previous {
			t.Fatal()
		}

		previous =
			string(entry.Key)
	}
}
