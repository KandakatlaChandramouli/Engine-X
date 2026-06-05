package core

import "testing"

func TestShiftSlotsRight(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	Insert(
		&p,
		[]byte("a"),
		[]byte("1"),
	)

	Insert(
		&p,
		[]byte("b"),
		[]byte("2"),
	)

	Insert(
		&p,
		[]byte("c"),
		[]byte("3"),
	)

	original :=
		ReadSlot(
			&p,
			1,
		)

	ShiftSlotsRight(
		&p,
		1,
	)

	shifted :=
		ReadSlot(
			&p,
			2,
		)

	if original.Offset != shifted.Offset {
		t.Fatal(
			"slot not shifted",
		)
	}
}
