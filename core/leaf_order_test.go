package core

import "testing"

func TestLeafSearchUnsortedInsert(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	SetPageType(
		&p,
		PageTypeLeaf,
	)

	Insert(
		&p,
		[]byte("zebra"),
		[]byte("1"),
	)

	Insert(
		&p,
		[]byte("apple"),
		[]byte("2"),
	)

	Insert(
		&p,
		[]byte("banana"),
		[]byte("3"),
	)

	idx :=
		LeafSearch(
			&p,
			[]byte("apple"),
		)

	if idx >= 0 {
		t.Fatalf(
			"binary search should fail on unsorted leaf",
		)
	}
}
