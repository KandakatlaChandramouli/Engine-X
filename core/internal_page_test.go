package core

import "testing"

func TestLeftChild(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	SetLeftChild(
		&p,
		123,
	)

	if LeftChild(
		&p,
	) != 123 {

		t.Fatal(
			"bad left child",
		)
	}
}
