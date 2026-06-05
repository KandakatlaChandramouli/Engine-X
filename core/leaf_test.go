package core

import "testing"

func TestLeafSearch(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		10,
	)

	SetPageType(
		&p,
		PageTypeLeaf,
	)

	Insert(
		&p,
		[]byte("apple"),
		[]byte("1"),
	)

	Insert(
		&p,
		[]byte("banana"),
		[]byte("2"),
	)

	Insert(
		&p,
		[]byte("cat"),
		[]byte("3"),
	)

	idx :=
		LeafSearch(
			&p,
			[]byte("banana"),
		)

	if idx != 1 {
		t.Fatalf(
			"expected 1 got %d",
			idx,
		)
	}
}
