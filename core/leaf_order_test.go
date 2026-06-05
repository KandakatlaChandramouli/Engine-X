package core

import "testing"

func TestLeafSearchRequiresSortedKeys(
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

	for i := uint16(0); i < PageSlotCount(&p); i++ {

		k,
			_,
			ok :=
			Get(
				&p,
				i,
			)

		if !ok {
			t.Fatal("corrupt slot")
		}

		t.Logf(
			"slot=%d key=%s",
			i,
			string(k),
		)
	}
}
