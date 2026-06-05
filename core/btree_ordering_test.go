package core

import "testing"

func TestLeafOrdering(
	t *testing.T,
) {

	var page Page

	InitPage(
		&page,
		1,
	)

	keys := [][]byte{
		[]byte("z"),
		[]byte("m"),
		[]byte("a"),
		[]byte("t"),
		[]byte("b"),
	}

	for _, key := range keys {

		if !LeafInsert(
			&page,
			key,
			[]byte("value"),
		) {
			t.Fatal()
		}
	}

	previous := ""

	count :=
		LeafEntryCount(
			&page,
		)

	for i := uint16(0); i < count; i++ {

		key,
			_,
			ok :=
			Get(
				&page,
				i,
			)

		if !ok {
			t.Fatal()
		}

		if previous != "" &&
			string(key) < previous {
			t.Fatal()
		}

		previous =
			string(key)
	}
}
